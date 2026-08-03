# Limitations

This section lists, without exception, what is missing or uncertain in the source material and in this repository.

## Missing from the original source document

1. Primary source not identified. The guide states it is based on a video that builds a Google Drive clone, but does not name the video, its author or channel, platform, or URL. Action needed: if you have this information, add it to references/references.bib before citing this repository publicly; until then, the video cannot be formally cited.
2. No formal bibliography. The original guide does not cite any paper, book, standard, or vendor documentation for the concepts it presents (load balancing algorithms, the cache-aside pattern, JWT-based authentication, etc.). These are widely known industry patterns, but no specific source is attributed to them in the material provided.
3. No empirical data or benchmarks. All numbers appearing in the guide (for example, 500 users opening the same file, a rate limit of 10 requests per minute, a 50 MB size limit, latency comparisons of about 20 ms versus about 1 second) are illustrative, pedagogical examples used to explain a concept. They are not the result of a measured experiment and must not be cited as empirical findings.
4. No author or institutional attribution beyond the repository owner. The study guide itself does not identify an author distinct from whoever compiled the notes.

## Limitations of this repository

1. Translation risk. The study guide was translated from Spanish to English by an AI assistant under human supervision. While concepts and structure were preserved carefully, subtle nuances of the original phrasing may not translate perfectly. Report any suspected translation issue via a GitHub issue.
2. Illustrative demos, not benchmarks. The three demonstrations in src/ (load balancing, cache-aside, rate limiting) are minimal, single-machine, Docker Compose setups meant to let a reader observe qualitative behavior. They have not been load tested, and no performance numbers should be inferred from running them casually.
3. No independent peer review. The technical content has been organized and cross-checked against the source guide and against publicly available system-design references (see references/consulted-repositories.md), but it has not undergone formal peer review.
4. Single documentary source. Because only one study guide was available, this repository cannot yet triangulate its claims against multiple independent primary sources.
5. Scope of research questions. The research questions in docs/02-research-questions.md are answered here at a conceptual, documentary level. They are not (yet) answered through controlled experiments; see docs/06-future-work.md for what that would require.

## What would need to change before treating this as a citable academic source

- Identification and citation of the original video (title, author or channel, platform, URL, access date).
- At least one round of independent technical review.
- If empirical claims are desired, a documented experimental setup (hardware, network conditions, load-testing tool, number of repetitions, statistical treatment) producing the data in results/.
