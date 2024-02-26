# clouduploader_cli

## blob storage uploader

Allows to upload a file to Azure Blob Storage.

needed:

- Go installed
- Azure Storage Account
- Azure Storage Container

how to use:

- git clone https://github.com/yourusername/azure-blob-storage-uploader.git
- cd azure-blob-storage-uploader
- Create environment file (.env) and add your Azure Storage account details:

  > - AZURE_STORAGE_ACCOUNT=storage_account_name
  > - AZURE_STORAGE_KEY=storage_account_key

- go build -o upload

./upload azblob://<container_name> example.png

- replace <container_name> with Azure Blob Storage container
- example.png with the file you want to upload.

### TODO:

- add progress bar or percentage upload completion
- enable file synchronization
- add encryption for security
