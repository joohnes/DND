<script lang="ts">
	import { writable } from 'svelte/store';
  	import { HOST } from "$lib/host";
	import Player from '../../components/Player.svelte';
	import { onMount, onDestroy } from 'svelte';

	let players = writable([]);
	var menuItem: HTMLElement
	onMount(async () => {
		menuItem = document.getElementById("menu-players")!
		if (menuItem != undefined) {
			menuItem.classList.remove("btn-ghost")
			menuItem.classList.remove("btn-outline")
			menuItem.classList.add("btn-primary")
		}
		
		const res = await fetch(HOST + 'players');
		const data = await res.json();
		players.set(data);
	});
	
	onDestroy(() => {
		if (menuItem != undefined) {
			menuItem.classList.add("btn-ghost")
			menuItem.classList.add("btn-outline")
			menuItem.classList.remove("btn-primary")
		}
	})

</script>
{#if $players != null}
	<div class="flex justify-center pb-6">
		<span class="text-3xl">Players</span>
	</div>
	<div class="flex flex-wrap gap-5 justify-center">
		{#each $players as player}
			<Player id={player} />
		{/each}
	</div>
	{:else}
	<div class="flex justify-center pb-6">
		<span class="text-3xl">No Players</span>
	</div>
{/if}

