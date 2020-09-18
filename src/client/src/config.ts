/** Describes the app configuration */
export interface Configuration {
  backend: {
    proto: string;
    host: string;
    port: number;
  };
}

/**
 * App Configuration
 */
export const config: Configuration = {
  backend: {
    proto: "http",
    port: 8082,
    host: "localhost",
  },
};
