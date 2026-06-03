// npm install firecrawl
import { Firecrawl } from 'firecrawl';

const app = new Firecrawl({ apiKey: "fc-74331d9dda7541979ce31ab52b22ec12"  });

// Scrape a website:
app.scrape('firecrawl.dev')
