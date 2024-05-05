// Mapped type для преобразования полей
export type CamelCaseKeys<T> = {
	[K in keyof T as Uncapitalize<string & K>]: T[K]
}
