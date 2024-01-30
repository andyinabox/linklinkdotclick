# 🖇 linkydink

This app is contained within a single binary in `dist/linkydink`

```bash
./dist/linkydink --port=8000 --host=127.0.0.1
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

 - [ ] Implement API endpoints with test data
   -  [x] `GET /api/links`
   -  [ ] `POST /api/links`
   -  [x] `GET /api/links/{id}`
   -  [ ] `DELETE /api/links/{id}`
   -  [ ] `PUT /api/links/{id}` 
 - [ ] Add persistence (sqlite?)
   - [ ] Create `linkservice`
 - [ ] Containerize
 - [ ] Deploy