import { config } from "../config";

/**
 * Formats URLs
 */
export class Url {
  constructor(
    public protocol: string,
    public host: string,
    public port: number
  ) {}

  /** url to a specific api endpoint */
  public api(endpoint: string): string {
    return `${this.protocol}://${this.host}:${this.port}/${endpoint}`;
  }
}

/**
 * The URL formatter
 */
export const url: Url = new Url(
  config.backend.proto,
  config.backend.host,
  config.backend.port
);
