# Resize an image (on the fly)
This AWS Lambda function renders an image, that is stored on a S3 media share. The funtion has the ability to resize the image or add a white border on the fly.

You can also build a standalone commandline application to resize all images inside a specified folder on the local filesystem.

Supports JPG and PNG
# Usage of the Lambda function
Submit an image path what is an subpath of the configured media share on S3. 

Add one of the following URL-Paramters to resize or add a border:  

| parameter | type   | mandatory | description                                   |
|-----------|--------|-----------|-----------------------------------------------|
| w         | number | no        | maximum width of the rendered image           |         
| h         | number | no        | maximum heigth of the rendered image          |
| b         | number | no        | white border to add around the rendered image |

# Usage of the standalone commandline application

Call the commandline application `resizeimage` with the following parameters: 

| parameter | type   | mandatory | example              | default                  | description                                   |
|-----------|--------|-----------|----------------------|--------------------------|-----------------------------------------------|
| -in       | string | no        | /home/cma/images     | .                        | input folder                                  |         
| -out      | string | no        | /home/cma/images/out | equal to the size string | output folder, will be created if necessary   |
| -size     | string | no        | 500x500              | 250x250                  | maximum deimension of the rendered image      |
| -border   | number | no        | 10                   | 0                        | white border to add around the rendered image |

# Build Parameter

You can customize the build with the following evironment variables:

| parameter        | type   | mandatory | default                         | description                                                                       |
|------------------|--------|-----------|---------------------------------|-----------------------------------------------------------------------------------|
| BASE_DOMAIN_NAME | string | no        | matzat.cloud                    | Base Domain for S3 Media Storage and S3 Lambda Repository                         |
| REPOSITORY_URL   | string | no        | lambda-repo.$(BASE_DOMAIN_NAME) | FQDN of the S3 Lambda Repository (defaults to a subdomain of the BASE_DOMAIN_NAME |
| REPOSITORY_NAME  | string | no        | resizeimages3                   | Name of the S3 Folder inside the S3 Lambda Repository                             |
