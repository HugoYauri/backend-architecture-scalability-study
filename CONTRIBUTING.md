# Contributing

This project follows the spirit of open, curated educational repositories such as system-design-primer and awesome-scalability: contributions are welcome, but every addition must be traceable and verifiable.

## Ground rules

1. No invented data or citations. If you add a claim, a benchmark, or a reference, it must be verifiable (a link, a paper, a reproducible experiment). If something is an assumption or illustrative example, label it explicitly as such.
2. 2. Keep the study guide and the research framing separate. Edits to docs/study-guide/*.md should stay faithful to the original source material; new analysis, opinions, or extensions belong in docs/06-future-work.md or a new docs/ file.
   3. 3. Demos must be reproducible. Any code added to src/ must include a docker-compose.yml (or equivalent) and a README.md explaining exactly how to run it and what it does or does not demonstrate.
      4. 4. Use clear, consistent naming. Follow the existing NN-kebab-case.md convention for docs and NN-kebab-case-demo/ for source folders.
        
         5. ## How to contribute
        
         6. 1. Fork the repository and create a branch: git checkout -b feature/short-description
            2. 2. Make your changes, following the ground rules above.
               3. 3. If you touched src/, run the demo locally and confirm it works from a clean docker compose up --build.
                  4. 4. Open a pull request describing what changed and why, and which files in docs/04-limitations.md are affected, if any.
                    
                     5. ## Reporting missing or incorrect information
                    
                     6. If you find a gap, for example you know the original video source, or you have access to real benchmark data, please open an issue instead of silently editing. This keeps the provenance of every claim auditable.
                     7. 
