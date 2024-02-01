# 🖇 linkydink

This app is contained within a single binary in `dist/linkydink`

```bash
./dist/linkydink --port=8000 --host=127.0.0.1 --dbfile=db/linkydink.db
```

## Build commands

```bash
# build single binary into dist/
make

# build and run server
make run

# run server, restart when files change
make watch
```

## Todo

 - [ ] Implement API endpoints with ~~test data~~ sqlite
   -  [x] `GET /api/links`
   -  [x] `POST /api/links`
   -  [x] `GET /api/links/{id}`
   -  [x] `DELETE /api/links/{id}`
   -  [x] `PUT /api/links/{id}` 
   -  [ ] ~~`PATCH /api/links/{id}`~~
 - [x] Add persistence (sqlite?)
   - [x] Create `linkrepository`
 - [ ] Functionality
   - [x] Fetching updated data for links
   - [x] Adding links
   - [ ] Update `LastClicked` after click
   - [ ] Re-ordering links on frontend
   - [ ] Deleting links
   - [ ] Editing link title
   - [ ] Ability to have non-RSS links
 - [ ] Containerize
 - [ ] Deploy
 - [ ] Add SSL
 - [ ] Authentication
