# HINAMe backend server  

## API endpoints  

URL prefix: http://api/\<RESOURCE\>

| RESOURCE | METHOD |PATH| ARGUMENTS | RESPONSE | DESCRIPTION |
|:--------:|:------:|:---|:----------|:---------|:------------|
|shelter  | GET    |/:id_max/:id_min   |  | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's list by id |
|         | GET    |/:max_latitude/:max_longitude/:min_latitude/:min_longitude | | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's list aroud user's geolocation |
|         | GET    |/name/:id_max/:id_min | | [{id: int, name: string, latitude: float, longitude: float, state: string}] | get shelter's name list from id list |
|board    | GET    |/body/:shelter_id   |  | html | get board information (HTML template) |
|         | GET    |/shelter_id/:updated_after   |  | [{id: int, shelter_id: int}] | get board list which recently updated |
|         | PUT    |/:id   | html | html | Update board html |
|comment  | GET    |/:board_id   |  | [{id: int, parent_id: int (nullable), title: string, body: string, author: string, updated_at: timestamp }] | get list of comments which specified by shelter_id |
|         | GET    |/reply/:parent_id |  | [{id: int, parent_id: int (nullable), title: string, body: string, author: string, updated_at: timestamp }] | get list of comments which specified by shelter_id |
|         | POST   |/   | {parent_id: int or null, title: string, body: string, author: string} |  | upload comment for specified board or comment |
|         | PUT    |/:id   | {title: string, body: string, author: string} |  | update a specified comment |
|         | DELETE |/:id   |  |  | delete a comment |

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

## Responsibility of each Layer  

| name | responsiblity |
|:------:|:-----|
| application layer | aggregate data and throw messages to presentation |
| domain layer | define all models |
| infrastructure layer | prepare interface for access to DB |
