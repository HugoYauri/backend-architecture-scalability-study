# Backend Architecture Scalability Study

> A structured, reproducible study of how backend architectures evolve as load and organizational complexity grow - from a single stateful server to a distributed, cached, rate-limited microservices system.
>
> ## English Summary
>
> This repository organizes and formalizes a case-study-based exploration of backend architecture evolution, originally compiled as a study guide derived from an instructional video that incrementally builds a "Google Drive clone" backend. The guide walks through eleven conceptual stages - from a single stateful server to a fully distributed system with an API gateway, authentication, asynchronous messaging, caching, a CDN, and rate limiting - with the recurring principle of separation of responsibilities applied at increasing scale.
>
> This project turns those notes into a research-style repository: it defines explicit research questions, a methodology combining literature-style synthesis with small, reproducible hands-on demonstrations (load balancing, cache-aside pattern, rate limiting), and clearly documents what is known from the source material versus what is missing (primary source citation, empirical benchmarks, peer-reviewed references) so nothing is presented as verified data unless it actually is.
>
> ## Resumen en Espanol
>
> Este repositorio organiza y formaliza una exploracion basada en un caso de estudio sobre la evolucion de arquitecturas de backend, originalmente compilada como una guia de estudio derivada de un video instructivo que construye progresivamente el backend de un "clon de Google Drive". La guia recorre once etapas conceptuales, desde un unico servidor con estado hasta un sistema distribuido completo con API gateway, autenticacion, mensajeria asincrona, cache, CDN y rate limiting, bajo el principio recurrente de separacion de responsabilidades aplicado a escalas cada vez mayores.
>
> Este proyecto convierte esas notas en un repositorio con enfoque de investigacion: define preguntas de investigacion explicitas, una metodologia que combina sintesis documental con demostraciones practicas reproducibles pequenas (balanceo de carga, patron cache-aside, rate limiting), y documenta con claridad que se conoce a partir del material fuente y que informacion falta (cita de la fuente original, datos empiricos, referencias revisadas por pares).
>
> ## Objectives
>
> 1. Formalize an informal study guide on backend architecture into a structured, citable, and auditable document set.
> 2. 2. Identify, for each architectural transition described in the source material, the problem it solves, the trade-off it introduces, and the conditions under which it applies.
>    3. 3. Provide small, reproducible, runnable demonstrations for the parts of the guide that are inherently technical/executable (load balancing, caching, rate limiting), without fabricating performance numbers that were not measured.
>       4. 4. Make explicit what information is missing from the source material (primary citation, formal references, empirical data) instead of silently filling gaps.
>         
>          5. ## Research Questions
>         
>          6. See docs/02-research-questions.md for the full list. Summary:
>         
>          7. - RQ1: Why does an architecture that works for a small number of concurrent users become the wrong choice at a much larger scale?
> - RQ2: Under what conditions does decoupling application state from the server actually deliver the claimed gains?
> - - RQ3: How do different load-balancing algorithms compare given heterogeneous request costs?
>   - - RQ4: What complexity does a microservices decomposition introduce, and when does it pay off?
>     - - RQ5: What is the qualitative difference between an in-memory cache and a CDN?
>       - - RQ6: How does rate limiting, built on the same caching infrastructure, protect a system?
>        
>         - ## Methodology
>        
>         - See docs/03-methodology.md. In summary: (1) documentary synthesis of the original study guide into docs/study-guide/, and (2) small reproducible hands-on demonstrations in src/ for the mechanisms that are inherently executable.
>        
>         - ## Repository Structure
>
> - docs/ - research framing and the translated/organized study guide
> - - src/ - minimal, runnable demonstrations (Docker Compose based)
>   - - data/ - placeholder for any raw/processed data from future experiments
>     - - results/ - placeholder for figures/metrics produced by running the demos
>       - - figures/ - architecture diagrams (Mermaid, text-based)
>         - - references/ - bibliography file and list of consulted repositories
>          
>           - ## Installation and Reproduction
>          
>           - ### Prerequisites
>          
>           - - Docker and Docker Compose v2
> - Go 1.21+ (only if running services outside Docker)
> - - Git
>  
>   - ### Clone
>  
>   - git clone https://github.com/HugoYauri/backend-architecture-scalability-study.git
>   - cd backend-architecture-scalability-study
>  
>   - ### Run the load-balancing demo (Blocks 3-4)
>
>   - cd src/01-load-balancing-demo
>   - docker compose up --build
>
>   - ### Run the cache-aside demo (Block 10)
>
>   - cd src/02-cache-aside-demo
>   - docker compose up --build
>
>   - ### Run the rate-limiting demo (Block 11)
>
>   - cd src/03-rate-limiting-demo
>   - docker compose up --build
>
>   - ### Tear down
>
>   - docker compose down -v
>
>   - ## Expected Results
>
>   - No empirical results currently exist in this repository. See docs/05-expected-results.md.
>
>   - ## Limitations
>
>   - See docs/04-limitations.md, including the missing primary source citation, the absence of formal references and empirical data in the original material, and the illustrative nature of the src/ demos.
>
>   - ## Future Work
>
>   - See docs/06-future-work.md.
>
>   - ## Consulted Repositories / Inspiration
>
>   - See references/consulted-repositories.md.
>
>   - ## Citation
>
>   - If you use this repository, please cite it using the metadata in CITATION.cff.
>
>   - ## License
>
>   - Released under the MIT License. See LICENSE.
>   - 
