<script lang="ts">
	import { onMount } from "svelte";
  import { HOST } from "$lib/host";

    /** @type {import('./$types').PageData} */
    export let data;
    const id = data.data
    const sendData = (e: any) => {
      let p: Player
      const formData = new FormData(e.target)
      p = {
        id: parseInt(id),
        name: formData.get("name")!.toString(),
        level: parseInt(formData.get("level")!.toString()),
        class: formData.get("class")!.toString(),
        race: formData.get("race")!.toString(),
        subrace: formData.get("subrace")!.toString(),
        strength: parseInt(formData.get("strength")!.toString()),
        endurance: parseInt(formData.get("endurance")!.toString()),
        perception: parseInt(formData.get("perception")!.toString()),
        intelligence: parseInt(formData.get("intelligence")!.toString()),
        agility: parseInt(formData.get("agility")!.toString()),
        accuracy: parseInt(formData.get("accuracy")!.toString()),
        charisma: parseInt(formData.get("charisma")!.toString())
      }

      fetch(HOST + "player/" + id, {
        method: "PUT",
        body: JSON.stringify(p)
      })
      window.location.href = window.location.origin + "/players"
    }
    let p: Player;
    onMount(async function () {
      const res = await fetch(HOST + 'player/' + id);
      const data = await res.json();
      p = {
        id: data.id,
        name: data.name,
        level: data.level,
        health: data.health,
        mana: data.mana,
        class: data.class,
        race: data.race,
        subrace: data.subrace,
        strength: data.strength,
        endurance: data.endurance,
        perception: data.perception,
        intelligence: data.intelligence,
        agility: data.agility,
        accuracy: data.accuracy,
        charisma: data.charisma,
        session: data.session,
        items: data.items
      };
    });
  </script>
{#if p != undefined}
<div class="h-screen my-auto items-center py-1 px-8 m-2 flex flex-col drop-shadow-xl">
  <form on:submit|preventDefault={sendData}>
      <fieldset class="border border-slate-900 rounded-lg p-5 shadow-2xl">
        <div class="flex gap-20">
          <div>
            <label for="name">Name</label>  
            <div>
             <input id="name" name="name" type="text" value={p.name} class="input input-bordered w-full max-w-md">
            </div>
      
            <label class="" for="level">Level</label>  
            <div>
              <input id="level" name="level" type="number" value={p.level} class="input input-bordered w-full max-w-xs">
            </div>
          
            <label class="" for="class">Class</label>
            <div>
              <input id="class" name="class" type="text" value={p.class} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="race">Race</label>
            <div >
              <input id="race" name="race" type="text" value={p.race} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="subrace">Subrace</label>
            <div >
              <input id="subrace" name="subrace" type="text" value={p.subrace} class="input input-bordered w-full max-w-xs">
            </div>
          </div>
          <div>
            <label  for="strength">Strength</label>  
            <div >
              <input id="strength" name="strength" type="number" value={p.strength} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="endurance">Endurance</label>  
            <div >
              <input id="endurance" name="endurance" type="number" value={p.endurance} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="perception">Perception</label>  
            <div >
              <input id="perception" name="perception" type="number" value={p.perception} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="intelligence">Intelligence</label>  
            <div >
            <input id="intelligence" name="intelligence" type="number" value={p.intelligence} class="input input-bordered w-full max-w-xs">
            </div>
          
            <label  for="agility">Agility</label>  
            <div >
            <input id="agility" name="agility" type="number" value={p.agility} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="accuracy">Accuracy</label>  
            <div >
            <input id="accuracy" name="accuracy" type="number" value={p.accuracy} class="input input-bordered w-full max-w-xs">
            </div>
          
          
            <label  for="charisma">Charisma</label>  
            <div >
            <input id="charisma" name="charisma" type="number" value={p.charisma} class="input input-bordered w-full max-w-xs">
            </div>
          </div>
        </div>
    <button class="btn"><input type="submit" value="Create"></button>
      </fieldset>
      </form>
    </div>  
  {/if}