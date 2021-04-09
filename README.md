# ladda
A utility written in Go that uploads files to your server _with **zero** dependencies_

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/willywill/ladda/test)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/willywill/ladda)
![GitHub Repo License](https://img.shields.io/github/license/willywill/ladda)

### Getting Started

Ladda runs on in a docker container on your webserver. It can respond to authorized requests to upload and store the files sent to it's endpoint on your webserver.

---

To install, pull the latest container:

```
docker pull ghcr.io/willywill/ladda:latest
```

Setup the `.env` file in the directory you plan to initialize the container in so that you at least have a `SECRET` defined. The application will not upload the file unless the `SECRET` matches the value passed in the `Authorization` header as a Base64 encoded value.

Run the container by running the `init-container.sh` file, and pass it the following optional values:

- `port` - the port exposed by your webserver
- `path` - the path to write the files to on the host webserver

By default, the port will map to `3001`, so if you use Nginx for example and have the following `sites-available` config, it should just work.

```
location / {
   proxy_pass http://localhost:3001;
}
```

By default, the file will write to a temp folder of the current directory that the script is initialized in.

---

To call the API:

- Send a POST request to `https://<your_domain>/api/v1/upload`
- Ensure that the `Content-Type` is set to some form of `multipart/form-data`
- Ensure that there is a `Authorization` header, that has a value of the `SECRET` env variable in Base64 format.
- By default, the max file size is 10MB, you can change this with the `MAX_FILE_SIZE` env variable in number of bytes.

### Upcoming Improvements

 - Configure allowable file MIME types
 - Configure API endpoint
