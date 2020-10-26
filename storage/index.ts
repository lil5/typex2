export type ArrayOne = string[]

export type Fans = Record<string, fan>

export type More = string

export interface Request {
	body: any
}

export interface RequestChild extends Request, fan {
	childID: string
}

interface fan {
	links: Record<string, string>
}

