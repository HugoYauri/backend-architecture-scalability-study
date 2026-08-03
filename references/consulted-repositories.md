# Consulted Repositories and Inspiration

Before designing this repository's structure, the following real, actively referenced GitHub repositories were reviewed for organizational, documentation, and reproducibility practices. No text or code was copied from them; only structural and procedural conventions were used as inspiration, as described below.

## donnemartin/system-design-primer

https://github.com/donnemartin/system-design-primer

What was taken as inspiration: organizing content by sub-topic with explicit trade-offs ("pros and cons") for each concept, a navigable index at the top of the document, and a closing "next steps" style study path. This directly informed the structure of docs/study-guide/ and docs/study-guide/12-synthesis-and-study-path.md.

## binhnguyennus/awesome-scalability

https://github.com/binhnguyennus/awesome-scalability

What was taken as inspiration: curating material by problem rather than by technology, linking out to primary sources instead of writing long paraphrased summaries, and having clear contribution guidelines. This informed CONTRIBUTING.md and the preference for short, attributable claims throughout docs/.

## drivendataorg/cookiecutter-data-science

https://github.com/drivendataorg/cookiecutter-data-science

What was taken as inspiration: a reproducible project layout that strictly separates data/, docs/, references/, and generated results/ or reports, plus a Makefile with convenience commands and a requirements.txt for environment reproduction. This directly informed the top-level layout of this repository (data/, results/, references/, Makefile, requirements.txt).

## microservices-patterns/ftgo-application

https://github.com/microservices-patterns/ftgo-application

What was taken as inspiration: pairing conceptual material (in that case, a book; here, the study guide) with runnable, Docker-based example code organized by the same chapters/blocks as the written material, so a reader can move between explanation and execution. This informed the src/ demos and their explicit links back to docs/study-guide/ files.

## Note on originality

This repository's actual content (the translated study guide, the research framing documents, and the demo source code) was written specifically for this project. The repositories above influenced only its organization and conventions, not its wording or code.
