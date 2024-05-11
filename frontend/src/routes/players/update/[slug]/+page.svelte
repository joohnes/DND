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

      fetch(HOST + "/player/" + id, {
        method: "PUT",
        body: JSON.stringify(p)
      })
      window.location.href = window.location.origin + "/players"
    }
    let p: Player;
    onMount(async function () {
		const res = await fetch(HOST + '/player/' + id);
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
<form class="form-horizontal" on:submit|preventDefault={sendData}>
  <fieldset>
  <div class="form-group">
    <label class="col-md-4 control-label" for="name">Name</label>  
    <div class="col-md-4">
    <input id="name" name="name" type="text" value={p.name} class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="level">Level</label>  
    <div class="col-md-4">
    <input id="level" name="level" type="number" value={p.level} class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="class">Class</label>
    <div class="col-md-4">
      <input id="class" name="class" type="text" value={p.class} class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="race">Race</label>
    <div class="col-md-4">
      <input id="race" name="race" type="text" value={p.race} class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="subrace">Subrace</label>
    <div class="col-md-4">
      <input id="subrace" name="subrace" type="text" value={p.subrace} class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="strength">Strength</label>  
    <div class="col-md-4">
    <input id="strength" name="strength" type="number" value={p.strength}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="endurance">Endurance</label>  
    <div class="col-md-4">
    <input id="endurance" name="endurance" type="number" value={p.endurance}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="perception">Perception</label>  
    <div class="col-md-4">
    <input id="perception" name="perception" type="number" value={p.perception}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="intelligence">Intelligence</label>  
    <div class="col-md-4">
    <input id="intelligence" name="intelligence" type="number" value={p.intelligence}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="agility">Agility</label>  
    <div class="col-md-4">
    <input id="agility" name="agility" type="number" value={p.agility}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="accuracy">Accuracy</label>  
    <div class="col-md-4">
    <input id="accuracy" name="accuracy" type="number" value={p.accuracy}  class="form-control input-md">
    </div>
  </div>
  
  <div class="form-group">
    <label class="col-md-4 control-label" for="charisma">Charisma</label>  
    <div class="col-md-4">
    <input id="charisma" name="charisma" type="number" value={p.charisma} class="form-control input-md">
    </div>
  </div>
  <input type="submit" value="Update">
  
  </fieldset>
  </form>
  {/if}