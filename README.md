### Setup Agent

Create a unique agent with CA from the server:

```
go run cmd/gen-agent/main.go --ca-cert <path_to_cert> --ca-key <path_to_key>
```

### Run Agent

After installing dependencies (bash, trivy, kopia), run the following:

```
sudo ./syntinel-agent-<uuid>
```

