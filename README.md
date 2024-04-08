# CRUD APi

Simple CRUD API wirtten in `Go` with firestore db.

## Requirements

This project use firestore emulator. You need to to set it up according to your env to make it work.
- install `firestore-cli`
- change `.env` and `.firebaserc` project id value to yours

## Documentation

This project use `swagger` as documentation. You can run this API and find it at http://localhost:8085/swagger/index.html

## Endpoints

### /api/v1

|**method**|**path**|**params**|**decription**|
|**method**|**path**|**params**|**decription**|
|`GET`|`health`|None|little health check|

### /api/v1/albums
|**method**|**path**|**params**|**decription**|
|**method**|**path**|**params**|**decription**|
|`GET`|`/albums`|None|List albums|
|`GET`|`/albums`|`:id`|Get album by id|
|`DELETE`|`/albums`|`:id`|Delete album by id|
|`PATCH`|`/albums`|`:id`|Update album by id|
|`POST`|`/albums`|None|Post album as JSON|

