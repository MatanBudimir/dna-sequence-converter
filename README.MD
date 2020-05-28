<h2>DNA Sequence Converter</h2>

Simple tool that converts DNA/RNA sequence to DNA, mRNA and amino acid sequence.

To host this website create `.env.local` file and fill it with your mysql credentials.

```.env
DATABASE=dna
USER=user
PASSWORD=password123
```

Once you've configured `.env`, run `go run migrate.go`. That will create and fill proteins table.

Once finished run `go run server.go` to start the application.
