<script lang="ts">
  const sendData = (e: any) => {
      let p: Player
      const formData = new FormData(e.target)
      p = {
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

      fetch(window.location.origin.replace("5173", "8000") + "/player", {
        method: "POST",
        body: JSON.stringify(p)
      })
      window.location.href = window.location.origin + "/players"
    }

</script>

<form class="form-horizontal" on:submit|preventDefault={sendData}>
    <fieldset>
    <div class="form-group">
      <label class="col-md-4 control-label" for="name">Name</label>  
      <div class="col-md-4">
      <input id="name" name="name" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="level">Level</label>  
      <div class="col-md-4">
      <input id="level" name="level" type="number" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="class">Class</label>
      <div class="col-md-4">
        <input id="class" name="class" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="race">Race</label>
      <div class="col-md-4">
        <input id="race" name="race" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="subrace">Subrace</label>
      <div class="col-md-4">
        <input id="subrace" name="subrace" type="text" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="strength">Strength</label>  
      <div class="col-md-4">
      <input id="strength" name="strength" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="endurance">Endurance</label>  
      <div class="col-md-4">
      <input id="endurance" name="endurance" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="perception">Perception</label>  
      <div class="col-md-4">
      <input id="perception" name="perception" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="intelligence">Intelligence</label>  
      <div class="col-md-4">
      <input id="intelligence" name="intelligence" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="agility">Agility</label>  
      <div class="col-md-4">
      <input id="agility" name="agility" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="accuracy">Accuracy</label>  
      <div class="col-md-4">
      <input id="accuracy" name="accuracy" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    
    <div class="form-group">
      <label class="col-md-4 control-label" for="charisma">Charisma</label>  
      <div class="col-md-4">
      <input id="charisma" name="charisma" type="number" placeholder="" class="form-control input-md">
      </div>
    </div>
    <input type="submit" value="Create">
    
    </fieldset>
    </form>
    