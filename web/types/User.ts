import * as z from "zod"

export const OAuthProviderSchema = z.object({
  username: z.string(),
  provider: z.string(),
})
export type OAuthProvider = z.infer<typeof OAuthProviderSchema>

export const UserSchema = z.object({
  id: z.string(),
  email: z.string(),
  name: z.string(),
  bio: z.string().nullable(),
  admin: z.boolean(),
  oauthProvider: OAuthProviderSchema.array().nullable(),
  createdAt: z.string(),
  updatedAt: z.string(),
  deletedAt: z.string().nullable(),
})
export type User = z.infer<typeof UserSchema>

