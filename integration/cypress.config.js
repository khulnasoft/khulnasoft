const { defineConfig } = require('cypress');

module.exports = defineConfig({
  video: false,
  defaultCommandTimeout: 20000,
  lighthouse: {
    performance: 90,
    accessibility: 100,
    'best-practices': 80,
    seo: 0,
    pwa: 0
  },
  e2e: {
    baseUrl: 'http://localhost:8080',
    specPattern: 'cypress/e2e/**/*.cy.js',
    supportFile: false,  // Disable the support file
  }
});
