# Social Network 01

## Description

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ultricies laoreet nisl. Duis sapien turpis, ultrices nec sem ut, rutrum rhoncus velit. Aliquam ac ullamcorper nunc, in varius ligula. Integer sodales tincidunt eleifend. Duis non molestie lacus. Donec porttitor lacus ut nulla condimentum, sed semper nibh sollicitudin. Sed eu neque non erat egestas lobortis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur velit ante, ornare at mi non, elementum commodo urna.

Each folder corresponds to a server:

-   [`frontend`](/frontend/) : a Next js frontend server.
-   [`backend`](/backend/) : a Golang backend API.

## Setup

### Frontend:

First you will need to setup the environment variables with a `.env.local` file with these two variables:

`frontend/.env.local`

```.env
NEXT_PUBLIC_API_URL=http://localhost:3001
```

Install dependencies with:

```console
cd frontend
npm i
```

Build the Next js App:

```console
npm run build
```

Start the Next js App:

```console
npm run start
```

---

### Backend:

Install dependencies with:

```console
cd backend
go mod download
```

Build the Golang Executable:

```console
go build -o backend
```

Run the migrations:

```console
./backend up
```

Start the API:

```console
./backend serve
```
