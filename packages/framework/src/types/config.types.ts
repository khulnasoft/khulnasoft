export type ClientOptions = {
  /**
   * Use Khulnasoft Cloud US (https://api.khulnasoft.co) or EU deployment (https://eu.api.khulnasoft.co). Defaults to US.
   */
  apiUrl?: string;

  /**
   * Specify your Khulnasoft secret key, to secure the Bridge Endpoint, and Khulnasoft API communication.
   * Khulnasoft communicates securely with your endpoint using a signed HMAC header,
   * ensuring that only trusted requests from Khulnasoft are actioned by your Bridge API.
   * The secret key is used to sign the HMAC header.
   */
  secretKey?: string;

  /**
   * Explicitly use HMAC signature verification.
   * Setting this to `false` will enable Khulnasoft to communicate with your Bridge API
   * without requiring a valid HMAC signature.
   * This is useful for local development and testing.
   *
   * In production you must specify an `secretKey` and set this to `true`.
   *
   * Defaults to true.
   */
  strictAuthentication?: boolean;
};
