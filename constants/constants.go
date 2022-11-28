package constants

// Context
const INVALID_CONTEXT = "middlewares/invalid-context"

// AUTHENTICATION CONSTANTS
const AUTH_UNAUTHORIZED = "auth/not-authorized"
const AUTH_NOT_AUTHENTICATED = "auth/not-authenticated"
const AUTH_PASSWORD_MISSMATCH = "auth/password-missmatch"
const AUTH_NOBODY_LOGGED = "auth/nobody-logged"
const AUTH_UNABLE_TO_HASH_PASSWORD = "auth/unable-to-hash-password"

const AUTH_MISSING_TOKEN = "auth/missing-token"
const AUTH_MISSING_ROLE = "auth/missing-role?role="
const AUTH_WRONG_TOKEN_FORMAT = "auth/wrong-token-format"
const AUTH_INVALID_TOKEN = "auth/invalid-token"
const AUTH_MISSING_PERMISSIONS = "auth/missing-permissions"
const AUTH_UNVERIFIED_EMAIL = "auth/unverified-email"
const AUTH_EMAIL_EXISTS = "auth/email-already-exists"

// MESSAGES CONTANTS
const UNABLE_TO_BIND_BODY = "controller/unable-to-bind-body"
const RESOURCE_EXISTS = "data-already-exists/resource="
const RESOURCE_NOT_FOUND = "data-not-found/resource="
const UNABLE_TO_DO_ACTION = "unable-to-do/action="
const SUCCESS_ACTION = "successfull/action="
const BAD_DATA = "bad-data/resource="
const BAD_PARAMS = "bad-params/param="
const MISSING_PARAM = "missing-params/param="
const DB_ERROR = "db-error"

// SERVER
const SERVER_ERROR = "server/bad-configuration"
