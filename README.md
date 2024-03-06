# clouduploader_cli

## blob storage uploader

Allows to upload a file to Azure Blob Storage.

needed:

- python installed
- Azure Storage Account
- Azure Storage Container

how to use:

- git clone https://github.com/yourusername/azure-blob-storage-uploader.git
- cd azure-blob-storage-uploader
- create environment file (.env) and add your Azure Storage account details:

  > - AZURE_ACCOUNT_NAME=storage_account_name
  > - AZURE_ACCOUNT_KEY=storage_account_key
  > - CONTAINER_NAME=storage_container_name

- pip install -r requirements.txt

python upload.py file1.txt file2.png ...

### TODO:

- add progress bar or percentage upload completion
- enable file synchronization
- add encryption for security
