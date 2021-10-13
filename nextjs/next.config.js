module.exports = {
  async rewrites() {
    // When running Next.js via Node.js (e.g. `dev` mode), proxy API requests
    // to the Go server.
    return [
      {
        source: "/v1/health",
        destination: "http://localhost:8080/v1/health",
      },
    ];
  },
  future: {
    webpack5: true,
  },
  trailingSlash: true,
};
