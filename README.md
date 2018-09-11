# HINAMe backend server  

## API endpoints  

URL prefix: http://api/\<RESOURCE\>

| RESOURCE | METHOD |PATH| ARGUMENTS | RESPONSE | DESCRIPTION |
|:--------:|:------:|:---|:----------|:---------|:------------|
|shelters  | GET    |/   | NONE | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's list |
|          | GET    |/   | max_lat: float, max_lon: float, min_lat: float, min_lon: float, exclude: [int] | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's list aroud user's geolocation |
|          | GET    |/names   | id: [int] | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's name list from id list |
|          | GET    |/details |latitude: float, longitude: float | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's list aroud user's geolocation |
|boards    | GET    |/body   | shelter_id | html | get board information (HTML template) |
|          | GET    |/shelter_id   | update_after: timestamp | [{id: int, shelter_id: int}] | get board list which recently updated |
|          | POST   |/   | id: int, md_text: string | html | Update board html |
|comments  | GET    |/   | shelter_id | [{id: int, parent_id: int (nullable), title: string, body: string, author: string, updated_at: timestamp }] | get list of comments which specified by shelter_id |
|          | GET    |/replies | shelter_id | [{id: int, parent_id: int (nullable), title: string, body: string, author: string, updated_at: timestamp }] | get list of comments which specified by shelter_id |
|          | POST   |/   | {parent_id: int|null, title: string, body: string, author: string} | NONE | upload comment for specified board or comment |
|          | POST   |/   | {id: int, title: string, body: string, author: string} | NONE | update a specified comment |
|          | DELETE |/   | id: int | NONE | delete a comment |

## DB schema  

### shelters

| COLUMN | TYPE | OPTION | DESCRIPTION |
|:------:|:-----|:-------|:------------|
| id     | int  | primary key, auto increment ||
| name   | varchar(256) | not null ||
| latitude   | float | not null ||
| longitude  | float | not null ||
| state  | enum('safe', 'filled', 'danger', 'unavailable') | not null, default='safe' ||

### boards
| COLUMN | TYPE | OPTION | DESCRIPTION |
|:------:|:-----|:-------|:------------|
| id     | int  | primary key, auto increment ||
| shelter_id  | int | not null, unique ||
| html   | text | not null, default='\<h3\>ようこそ\</h3\>' ||
| update_at  | timestamp | not null, default=now() on update now() ||

### commments
| COLUMN | TYPE | OPTION | DESCRIPTION |
|:------:|:-----|:-------|:------------|
| id     | int  | primary key, auto increment ||
| board_id  | int | not null ||
| parent_id  | int | ||
| title  | varchar(256) | not null,  default='タイトルなし' ||
| author | varchar(256) | not null,  default='住民' ||
| body   | text | not null ||
| updated_at   | timestamp | not null, default=now() on update now() ||
