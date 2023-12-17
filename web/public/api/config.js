import { useSession } from "../hooks/session.js";

// default server address.
export const instance = (process && process.env.SERVER) || "";
