export const snakeCaseData = (obj: Object) => {
	console.log(obj)

	return Object.fromEntries(
		Object.entries(obj).map(([key, value]) => [
			key.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`),
			value,
		])
	)
}

export const camelCaseData = (obj: Object) => {
	return Object.fromEntries(
		Object.entries(obj).map(([key, value]) => {
			return [
				key.replace(
					/_[a-z]/g,
					match => `${match.replace('_', '').toUpperCase()}`
				),
				value,
			]
		})
	)
}

export const getNumberInStr = (str?: string): string | null => {
	if (!str) return null
	const regExp = /\d+/g
	const result = regExp.exec(str)
	return result?.length ? result[0] : null
}
