// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}

	interface Item {
		id?: number;
		name: string;
		description: string;
		ability: string;
		rarity: string;
		strength: number;
		endurance: number;
		perception: number;
		intelligence: number;
		agility: number;
		accuracy: number;
		charisma: number;
		quantity: number;
		attack: number;
		defense: number;
		permille: number;
		slot?: number;
		owner?: string;
	}

	interface Player {
		id?: number;
		name: string;
		level: number;
		health?: number;
		mana?: number;
		class: string;
		race: string;
		subrace: string;
		strength: number;
		endurance: number;
		perception: number;
		intelligence: number;
		agility: number;
		accuracy: number;
		charisma: number;
		session?: number;
		items?: number[];
		equipped?: number[];
		alcohol_level?: number;
		zgon?: boolean;
	}

	interface PlayerResponse {
		id?: number;
		name: string;
		level: number;
		health?: number;
		mana?: number;
		class: string;
		race: string;
		subrace: string;
		strength: string;
		endurance: string;
		perception: string;
		intelligence: string;
		agility: string;
		accuracy: string;
		charisma: string;
		session?: number;
		items?: number[];
		equipped?: number[];
		alcohol_level?: number;
		zgon?: boolean;
	}

	interface Bag {
		owner: number;
		items: number[];
	}
}

export {};
