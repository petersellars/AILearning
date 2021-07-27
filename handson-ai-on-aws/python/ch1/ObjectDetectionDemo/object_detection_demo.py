from storage_service import StorageService
from recognition_service import RecognitionService

import argparse

parser = argparse.ArgumentParser()
parser.add_argument('--bucket', required=True, help='foo help')
args = parser.parse_args()

bucket_name = ''
if args.bucket:
    bucket_name = args.bucket

storage_service = StorageService()
recognition_service = RecognitionService()

for file in storage_service.get_all_files(bucket_name):
    if file.key.endswith('.jpg'):
        print('Objects detected in image ' + file.key + ':')
        labels = recognition_service.detect_objects(file.bucket_name, file.key)
        for label in labels:
            print('--' + label['Name'] + ': ' + str(label['Confidence']))