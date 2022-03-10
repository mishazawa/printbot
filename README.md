# printbot

It prints jpeg pictures and obj files in minecraft world using RCON protocol.

<img src="data/demo.png"/>

### Usage

Create `.env` file

```sh
# .env file

RCON_HOST=
RCON_PORT=
RCON_PASSWORD=
```

Run script

```sh

go run . -file=<path to image or obj file> -x=<x> -y=<y> -z=<z> [-remove] [-host=] [-port=] [-password=]

```

### etc

Mesh should have a lot of vertices to be proper filled by blocks. Olso it should be above ground (y > 0).

Only jpg and obj files.
