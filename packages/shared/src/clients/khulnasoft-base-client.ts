// Base HttpError class with response field
export class HttpError extends Error {
  constructor(
    public responseText: string,
    public status: number,
    public response: Response // Add response field
  ) {
    super(`${status}: ${responseText}`);
    this.name = this.constructor.name;
  }

  override toString(): string {
    return `${this.name} (status: ${this.status}): ${this.responseText}`;
  }
}

// Specific error classes extending HttpError
export class KhulnasoftBadRequestError extends HttpError {}
export class KhulnasoftUnauthorizedError extends HttpError {}
export class KhulnasoftForbiddenError extends HttpError {}
export class KhulnasoftNotFoundError extends HttpError {}
export class KhulnasoftInternalServerError extends HttpError {}
export class KhulnasoftNotImplementedError extends HttpError {}
export class KhulnasoftBadGatewayError extends HttpError {}
export class KhulnasoftServiceUnavailableError extends HttpError {}
export class KhulnasoftGatewayTimeoutError extends HttpError {}
export class KhulnasoftRedirectError extends HttpError {
  redirectUrl: string;

  constructor(responseText: string, status: number, redirectUrl: string, response: Response) {
    super(responseText, status, response); // Pass response to the base class
    this.redirectUrl = redirectUrl;
  }
}

// Map of status codes to specific error classes
const errorMap: Record<number, new (responseText: string, status: number, response: Response) => HttpError> = {
  400: KhulnasoftBadRequestError,
  401: KhulnasoftUnauthorizedError,
  403: KhulnasoftForbiddenError,
  404: KhulnasoftNotFoundError,
  500: KhulnasoftInternalServerError,
  501: KhulnasoftNotImplementedError,
  502: KhulnasoftBadGatewayError,
  503: KhulnasoftServiceUnavailableError,
  504: KhulnasoftGatewayTimeoutError,
};

// Type for the fetch function
type FetchFunction = () => Promise<Response>;

// Result class for handling success and failure
export class KhulnasoftRestResult<T, E> {
  public isSuccess: boolean;
  public value?: T;
  public error?: E;

  private constructor(isSuccess: boolean, value?: T, error?: E) {
    this.isSuccess = isSuccess;
    this.value = value;
    this.error = error;
  }

  static ok<T>(value: T): KhulnasoftRestResult<T, never> {
    return new KhulnasoftRestResult<T, never>(true, value);
  }

  static fail<E>(error: E): KhulnasoftRestResult<never, E> {
    return new KhulnasoftRestResult<never, E>(false, undefined as never, error);
  }

  public isSuccessResult(): this is { value: T; error: never } {
    return this.isSuccess;
  }
}

// Functional version of KhulnasoftBaseClient
export const createKhulnasoftBaseClient = (baseUrl: string, headers: HeadersInit = {}) => {
  const defaultHeaders = {
    'Content-Type': 'application/json',
    ...headers,
  };

  const buildUrl = (endpoint: string): string => `${baseUrl}${endpoint}`;

  const safeFetch = async <T>(url: string, fetchFunc: FetchFunction): Promise<KhulnasoftRestResult<T, HttpError>> => {
    const response: Response = await fetchFunc();

    if (response.ok) {
      const jsonData = await response.json();

      return KhulnasoftRestResult.ok(jsonData.data);
    }

    if (response.status >= 300 && response.status < 400) {
      const responseText = await getErrorResponse(response);
      const redirectError = new KhulnasoftRedirectError(
        responseText,
        response.status,
        response.headers.get('Location') || '',
        response // Pass the response object
      );

      return KhulnasoftRestResult.fail(redirectError);
    }

    const ErrorClass = errorMap[response.status] || HttpError;
    const responseText = await getErrorResponse(response);
    const errorResult = new ErrorClass(responseText, response.status, response); // Pass the response object

    return KhulnasoftRestResult.fail(errorResult);
  };

  async function getErrorResponse(response: Response): Promise<string> {
    // Try to parse the response as JSON
    try {
      const json = await response.json();

      return JSON.stringify(json); // Return the JSON as a string
    } catch {
      // If JSON parsing fails, fallback to text response
      return await response.text();
    }
  }

  const safeGet = async <T>(endpoint: string): Promise<KhulnasoftRestResult<T, HttpError>> => {
    return await safeFetch(endpoint, () =>
      fetch(buildUrl(endpoint), {
        method: 'GET',
        headers: defaultHeaders,
      })
    );
  };

  const safePut = async <T>(endpoint: string, data: object): Promise<KhulnasoftRestResult<T, HttpError>> => {
    return await safeFetch(endpoint, () =>
      fetch(buildUrl(endpoint), {
        method: 'PUT',
        headers: defaultHeaders,
        body: JSON.stringify(data),
      })
    );
  };

  const safePost = async <T>(endpoint: string, data: object): Promise<KhulnasoftRestResult<T, HttpError>> => {
    return await safeFetch(endpoint, () =>
      fetch(buildUrl(endpoint), {
        method: 'POST',
        headers: defaultHeaders,
        body: JSON.stringify(data),
      })
    );
  };
  const safePatch = async <T>(endpoint: string, data: object): Promise<KhulnasoftRestResult<T, HttpError>> => {
    return await safeFetch(endpoint, () =>
      fetch(buildUrl(endpoint), {
        method: 'PATCH',
        headers: defaultHeaders,
        body: JSON.stringify(data),
      })
    );
  };

  const safeDelete = async (endpoint: string): Promise<KhulnasoftRestResult<void, HttpError>> => {
    return await safeFetch(endpoint, () =>
      fetch(buildUrl(endpoint), {
        method: 'DELETE',
        headers: defaultHeaders,
      })
    );
  };

  return {
    safeGet,
    safePut,
    safePost,
    safeDelete,
    safePatch,
  };
};
