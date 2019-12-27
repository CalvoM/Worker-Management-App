# Workers Maintenance App
A simple app for maintaining workers, with Go backend.

## Getting started
Run this command below to install the npm modules needed in the project.
```js
npm install
```
Run this command in the root folder to get the project started.
```js
npm run server
```
## Note
If you run into **CORS-HEADER error**, go into the *cors-anywhere* module located here, relative to the root folder,
```sh
/node_modules/cors-anywhere/
```
Change the port number if needed in the *server.js* file and then run
```js
npm run start
```
To run the server code, go into the *server* folder and run
```sh
./db
```
Make sure to enter the correct configuration for the postgres DB in the **db.go** file.