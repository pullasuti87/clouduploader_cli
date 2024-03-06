import os
from dotenv import load_dotenv
from azure.storage.blob import BlobServiceClient
import sys

load_dotenv()


account_name = os.getenv("AZURE_ACCOUNT_NAME")
account_key = os.getenv("AZURE_ACCOUNT_KEY")
container_name = os.getenv("CONTAINER_NAME")


files = sys.argv[1:]


def upload():
    blob_client = BlobServiceClient(
        account_url=f"https://{account_name}.blob.core.windows.net",
        credential=account_key,
    )
    container_client = blob_client.get_container_client(container_name)

    for i in files:
        blob_name = os.path.basename(i)
        blob_client = container_client.get_blob_client(blob_name)

        with open(i, "rb") as data:
            blob_client.upload_blob(data)

        print(f"file '{i}' uploaded to blob storage!")


if __name__ == "__main__":

    if len(sys.argv) < 2:
        print("use: python upload.py <file1> <file2> ...")
        sys.exit(1)

    upload()
