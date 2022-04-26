build_dir := build

name_cmd := resizeimage
name_lambda := resizeimages3

BASE_DOMAIN_NAME?=matzat.cloud
REPOSITORY_URL?=lambda-repo.$(BASE_DOMAIN_NAME)
REPOSITORY_NAME?=$(name_lambda)

pkg_base := bitbucket.org/christian-m/$(name_cmd)
pkg_cmd := bitbucket.org/christian-m/$(name_cmd)/cmd/resizeimage
pkg_lambda := bitbucket.org/christian-m/$(name_cmd)/cmd/lambda

version := $(if $(shell git describe --tags --abbrev=0),$(shell git describe --tags --abbrev=0),build_$(shell git rev-parse --short HEAD))

build = GOOS=$(1) GOARCH=$(2) go build -o $(build_dir)/$(3)$(4) $(5)
tar = mkdir -p $(build_dir)/cli && cd $(build_dir)/cli && tar -cvzf $(1)_$(version).tar.gz ../$(2)$(3) && rm ../$(2)$(3)
zip = mkdir -p $(build_dir)/cli && cd $(build_dir)/cli && zip $(1)_$(version).zip ../$(2)$(3) && rm ../$(2)$(3)
lambda = mkdir -p $(build_dir)/lambda && cd $(build_dir)/lambda && zip $(1).zip ../$(2)$(3) && rm ../$(2)$(3)

.PHONY: all windows macos linux lambda deploy clean

all: macos linux windows

clean:
	rm -rf $(build_dir)/

dep:
	go get $(pkg_base)/...

fmt:
	go fmt $(pkg_base)/...

test: dep fmt
	go test -v $(pkg_base)/...

##### LINUX BUILDS #####
linux: test build/linux_arm.tar.gz build/linux_arm64.tar.gz build/linux_amd64.tar.gz

build/linux_amd64.tar.gz:
	$(call build,linux,amd64,$(name_cmd),,$(pkg_cmd))
	$(call tar,linux_amd64,$(name_cmd),)

build/linux_arm.tar.gz:
	$(call build,linux,arm,$(name_cmd),,$(pkg_cmd))
	$(call tar,linux_arm,$(name_cmd),)

build/linux_arm64.tar.gz:
	$(call build,linux,arm64,$(name_cmd),,$(pkg_cmd))
	$(call tar,linux_arm64,$(name_cmd),)

##### MACOS BUILDS #####
macos: test build/macos_amd64.tar.gz build/macos_arm64.tar.gz

build/macos_amd64.tar.gz:
	$(call build,darwin,amd64,$(name_cmd),,$(pkg_cmd))
	$(call tar,darwin_amd64,$(name_cmd),)

build/macos_arm64.tar.gz:
	$(call build,darwin,arm64,$(name_cmd),,$(pkg_cmd))
	$(call tar,darwin_arm64,$(name_cmd),)

##### WINDOWS BUILDS #####
windows: test build/windows_amd64.zip

build/windows_amd64.zip:
	$(call build,windows,amd64,$(name_cmd),.exe,$(pkg_cmd))
	$(call zip,windows_amd64,$(name_cmd),.exe)

##### AWS-LAMBDA BUILDS (LINUX AMD64) #####
lambda: build/lambda_linux_amd64.zip

build/lambda_linux_amd64.zip:
	$(call build,linux,amd64,$(name_lambda),,$(pkg_lambda))
	$(call lambda,$(name_lambda),$(name_lambda),)

deploy-dev: lambda
	cd $(build_dir) && aws s3 cp $(name_lambda)_$(version).zip s3://$(REPOSITORY_URL)/dev-$(REPOSITORY_NAME)/$(version)/dev-$(name_lambda).zip --content-type application/zip
	cd deployments/dev && terraform apply --auto-approve -var lambda_name=$(name_lambda) -var lambda_version=$(version) -var base_domain_name=$(BASE_DOMAIN_NAME)

deploy-prod: lambda
	cd $(build_dir) && aws s3 cp $(name_lambda)_$(version).zip s3://$(REPOSITORY_URL)/$(REPOSITORY_NAME)/$(version)/$(name_lambda).zip --content-type application/zip
	cd deployments/prod && terraform apply --auto-approve -var lambda_name=$(name_lambda) -var lambda_version=$(version) -var base_domain_name=$(BASE_DOMAIN_NAME)

destroy:
	cd deployments && terraform destroy --auto-approve