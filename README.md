# The Anarchy Solo Challenge Generator
I got the game for the holidays in 2025 and I really love it. I noticed that there hasn't been a solo challenge set up (maybe because it would require a lot of spoiler tags/information), so I decided to make a program that would generate everything.

# How it works
The `components.toml` file has a list of all the components that are necessary for generating the solo challenge. Since the domain cards are hidden (and not uniquely identifiable), I need to generate a complete deck. The below table shows how many possible domain cards need to exist. They probably wouldn't all be used, but this program generates them all.

102 cards - 24 cards per shuffle

| Area                     | Num of Cards |
| ------------------------ | ------------ |
| Knight's Training        | 18           |
| St. Valentine's Festival | 12           |
| Tournaments              | 36           |
| Brewhouse                | 12           |
| Michaelmas               | 12           |
| Lammas                   | 12           |

## Verifying releases

Releases are built in GitHub Actions from tagged source and include:

- platform binaries
- `checksums.txt`
- GitHub artifact attestations for build provenance

Pushing a tag such as `v1.2.3` triggers the release workflow. The binaries should be verified before extracting or running them.

To verify a downloaded release asset against GitHub's signed build provenance:

```bash
gh attestation verify ./<downloaded-release-asset> --repo nikkon226/anarchy-solo-challenge
```

To verify the checksum file locally after downloading `checksums.txt`:

```bash
sha256sum -c checksums.txt
```

If the repository has GitHub immutable releases enabled in repository settings, users can also verify the published release and that a local file matches a release asset:

```bash
gh release verify <tag> --repo nikkon226/anarchy-solo-challenge
gh release verify-asset <tag> ./<downloaded-release-asset> --repo nikkon226/anarchy-solo-challenge
```

This setup provides GitHub-signed provenance for release artifacts. It does not provide platform-native executable signing such as Apple code signing/notarization or Windows Authenticode.
