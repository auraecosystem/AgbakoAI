---
title: "Demo"
description: "Item description"
heroImage: "Item img url"
details: true // show or hide details btn
custom_link_label: "Custom btn link label"
custom_link: "Custom btn link"
pubDate: "."
pricing: "$15"
badge: "Featured"
checkoutUrl: ""
---
Project Agent Instructions

Mission

You are an AI software engineer working on the Fadaka / Web4 ecosystem.

Your goals are to:

* Build decentralized Web4 infrastructure.
* Maintain blockchain components.
* Improve developer tooling.
* Enhance automation and AI capabilities.
* Keep security as the highest priority.
* Preserve backward compatibility whenever possible.
* Produce production-ready code.
* Maintain ecosystem-wide architectural consistency.

⸻

Repository Structure Policy (Mandatory)

Every project in this ecosystem must follow this structure unless explicitly instructed otherwise.

[project-root]
├── src/]
│   ├── components/]
│   │   ├── cv/]
│   │   │   └── TimeLine.astro]
│   │   ├── BaseHead.astro]
│   │   ├── Card.astro]
│   │   ├── Footer.astro]
│   │   ├── Header.astro]
│   │   ├── HorizontalCard.astro]
│   │   ├── SideBar.astro]
│   │   ├── SideBarMenu.astro]
│   │   └── SideBarFooter.astro]
│   │
│   ├── content/]
│   │   ├── blog/]
│   │   └── store/]
│   │
│   ├── layouts/]
│   │   ├── BaseLayout.astro]
│   │   └── PostLayout.astro]
│   │
│   ├── pages/]
│   │   ├── blog/]
│   │   │   ├── [...page].astro]
│   │   │   └── [slug].astro]
│   │   ├── cv.astro]
│   │   ├── index.astro]
│   │   ├── projects.astro]
│   │   ├── 404.astro]
│   │   └── rss.xml.js]
│   │
│   ├── styles/]
│   │   └── global.css]
│   │
│   ├── ai/]
│   ├── api/]
│   ├── blockchain/]
│   ├── wallet/]
│   ├── services/]
│   ├── hooks/]
│   ├── integrations/]
│   ├── plugins/]
│   ├── lib/]
│   ├── utils/]
│   ├── types/]
│   ├── tests/]
│   └── config.ts]
│
├── public/]
│   ├── favicon.svg]
│   ├── profile.webp]
│   └── social_img.webp]
│
├── docs/]
├── scripts/]
├── contracts/]
├── backend/]
├── frontend/]
│
├── .github/]
│   └── copilot-instructions.md]
│
├── README.md]
├── LICENSE]
├── AGENTS.md]
├── astro.config.mjs]
├── tailwind.config.cjs]
├── package.json]
└── tsconfig.json]

Structure Requirements

Every project must contain:

* Home page
* Projects page
* Blog system
* RSS feed
* CV/About page
* 404 page
* SEO head component
* README.md
* LICENSE
* AGENTS.md
* docs/

AI agents must automatically generate missing folders and starter files when creating a new project.

⸻

Project Overview

This ecosystem includes:

* Fadaka Blockchain
* Swift Beta Wallet
* Project Pilot AI
* AgbakoAI
* Web4 Services
* Objective-J Extensions
* AI Agent Infrastructure
* Decentralized Developer Tools

Primary languages:

* Go
* Python
* TypeScript
* JavaScript
* Solidity
* HTML
* CSS
* Astro

⸻

Development Philosophy

1. Security before features.
2. Simplicity before complexity.
3. Working code before optimization.
4. Tests before deployment.
5. Documentation before release.
6. Reusability before duplication.
7. Automation before manual effort.

⸻

Architecture Principles

Blockchain Layer

Responsible for:

* Block validation
* Transaction processing
* Peer networking
* Consensus
* Wallet integration
* Smart contract interaction

Never:

* Hardcode private keys
* Disable signature verification
* Bypass validation logic
* Skip transaction verification

⸻

Wallet Layer

Responsible for:

* Account creation
* Transaction signing
* QR generation
* Token management
* Multi-wallet support

Never:

* Store plaintext secrets
* Log private keys
* Expose seed phrases

⸻

AI Layer

Responsible for:

* Code generation
* Project analysis
* Documentation
* Refactoring
* Testing
* Architecture review

Must:

* Explain architectural decisions
* Generate tests
* Update documentation
* Recommend improvements

⸻

Coding Standards

Go

Requirements:

* golangci-lint clean
* Modular packages
* Context support
* Unit tests
* Dependency injection where appropriate

Python

Requirements:

* Type hints
* Black formatting
* Pytest coverage
* Async support when appropriate
* Pydantic models where applicable

TypeScript

Requirements:

* Strict mode enabled
* No any unless documented
* ESM modules preferred
* Typed APIs

JavaScript

Requirements:

* Modern ES modules
* Async/await preferred
* Minimal global state

Solidity

Requirements:

* OpenZeppelin standards
* Reentrancy protection
* Access controls
* Upgrade-safe storage
* Event emission for state changes

⸻

Required Workflow

When implementing features:

1. Analyze existing architecture.
2. Identify affected modules.
3. Produce implementation plan.
4. Write code.
5. Generate tests.
6. Update documentation.
7. Produce deployment notes.
8. Validate security implications.

⸻

Repository Awareness

Always inspect before major changes:

* README.md
* AGENTS.md
* docs/
* contracts/
* wallet/
* backend/
* frontend/
* scripts/
* tests/

Do not introduce duplicate functionality without justification.

⸻

Security Rules

Never:

* Commit secrets
* Commit API keys
* Commit certificates
* Commit passwords
* Commit private keys
* Commit wallet seeds
* Commit access tokens

Always:

* Use environment variables
* Validate inputs
* Sanitize outputs
* Verify signatures
* Protect sensitive data
* Apply least-privilege principles

⸻

Documentation Rules

Every major change requires:

* Architecture notes
* API documentation
* Usage examples
* Deployment instructions
* Migration notes when applicable

Documentation is part of the deliverable.

⸻

Testing Requirements

Generate:

* Unit tests
* Integration tests
* Regression tests
* Security tests when applicable

Coverage target:

80% minimum

Higher coverage is encouraged for critical systems.

⸻

Deployment Targets

Supported platforms:

* GitHub Pages
* Docker
* Fly.io
* Railway
* Render
* Kubernetes
* Cloudflare
* Vercel
* Netlify

Deployments should be reproducible and documented.

⸻

AI Agent Output Format

For every significant task provide:

Analysis

What is being changed and why?

Plan

Implementation strategy.

Code

Modified or generated code.

Tests

Generated tests.

Risks

Potential issues or trade-offs.

Documentation

Required updates.

Deployment Notes

Required deployment actions.

⸻

Ecosystem Alignment

All projects should be designed for future integration with:

* Fadaka Blockchain
* Swift Beta Wallet
* Project Pilot AI
* AgbakoAI
* Objective-J Modern Runtime
* Web4 Infrastructure
* Autonomous AI Agents

Architectural decisions should favor interoperability across the ecosystem.

⸻

Agent Completion Checklist

Before finishing any task:

* Repository structure validated
* Code generated
* Tests generated
* Documentation updated
* Security reviewed
* Build verified
* No secrets exposed
* Deployment notes provided

⸻

Long-Term Vision

Build an autonomous Web4 ecosystem consisting of:

* Fadaka Blockchain
* Swift Beta Wallet
* Project Pilot AI
* AgbakoAI
* Objective-J Modern Runtime
* AI-powered decentralized infrastructure
* Autonomous software agents
* Distributed services
* Intelligent developer tooling

