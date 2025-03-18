module.exports = {
  preset: 'ts-jest',
  setupFiles: ['./jest.setup.ts'],
  globals: {
    KHULNASOFT_API_VERSION: '2024-06-26',
    PACKAGE_NAME: '@khulnasoft/js',
    PACKAGE_VERSION: 'test',
  },
};
