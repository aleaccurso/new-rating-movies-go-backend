package middlewares

// SERVER
const SERVER_ERROR = "server/bad-configuration"

// AUTHENTICATION
const AUTH_MISSING_TOKEN = "auth/missing-token"
const AUTH_MISSING_ROLE = "auth/missing-role?role="
const AUTH_WRONG_TOKEN_FORMAT = "auth/wrong-token-format"
const AUTH_INVALID_TOKEN = "auth/invalid-token"
const AUTH_MISSING_PERMISSIONS = "auth/missing-permissions"
const AUTH_UNVERIFIED_EMAIL = "auth/unverified-email"

// Context
const INVALID_CONTEXT = "middlewares/invalid-context"
