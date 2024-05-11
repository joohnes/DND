<script lang="ts">
    import { onMount } from 'svelte';
	import Modal from "./Modal.svelte"
  	import { HOST } from "$lib/host";

	let i: Item;
	onMount(async function () {
		const res = await fetch(HOST + 'item/' + id);
		const data = await res.json();
		i = {
			id: data.id,
			name: data.name,
			description: data.description,
			ability: data.ability,
			rarity: data.rarity,
			strength: data.strength,
			endurance: data.endurance,
			perception: data.perception,
			intelligence: data.intelligence,
			agility: data.agility,
			accuracy: data.accuracy,
			charisma: data.charisma,
			quantity: data.quantity
		};
	});

	export let id: number = 0;
</script>

<div>
    {#if i != undefined}
    <a href={window.location.origin+"/players/item/" +i.id}><div class="img">{i.name}</div></a>
    {/if}
</div>

<style>
    .img {
        width: 20px;
        height: 20px;
    }
</style>