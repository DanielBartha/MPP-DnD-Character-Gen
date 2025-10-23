let allCharacters = [];

window.addEventListener("DOMContentLoaded", async () => {
  await loadCharacters();

  const select = document.getElementById("characterSelect");
  select.addEventListener("change", () => {
    const name = select.value;
    const character = allCharacters.find(c => c.Name === name);
    if (character) fillSheet(character);
  });
});

async function loadCharacters() {
  try {
    const res = await fetch("/data/settings.json");
    if (!res.ok) throw new Error("Failed to load characters JSON file");

    allCharacters = await res.json();

    const select = document.getElementById("characterSelect");
    select.innerHTML = '<option value="">-- Choose character --</option>';

    allCharacters
      .sort((a, b) => a.Name.localeCompare(b.Name))
      .forEach(ch => {
        const opt = document.createElement("option");
        opt.value = ch.Name;
        opt.textContent = ch.Name;
        select.appendChild(opt);
      });

  } catch (err) {
    console.error(err);
    alert("Could not load character data. Run `go run main.go serve`");
  }
}

function fillSheet(ch) {
  setVal("charname", ch.Name);
  setVal("classlevel", `${ch.Class} ${ch.Level}`);
  setVal("background", ch.Background);
  setVal("race", ch.Race);
  setVal("proficiencybonus", `+${ch.Proficiency ?? 0}`);

  setStat("Strength", ch.Stats?.Str, ch.Stats?.StrMod);
  setStat("Dexterity", ch.Stats?.Dex, ch.Stats?.DexMod);
  setStat("Constitution", ch.Stats?.Con, ch.Stats?.ConMod);
  setStat("Intelligence", ch.Stats?.Intel, ch.Stats?.IntelMod);
  setStat("Wisdom", ch.Stats?.Wis, ch.Stats?.WisMod);
  setStat("Charisma", ch.Stats?.Cha, ch.Stats?.ChaMod);

  setVal("ac", `${ch.ArmorClass}`);
  setVal("initiative", `${ch.InitiativeBonus}`);
  setVal("passiveperception", `${ch.PassivePerception}`);

  const eq = [];
  if (ch.Equipment) {
    if (ch.Equipment.Weapon) {
      const main = ch.Equipment.Weapon["main hand"];
      const off = ch.Equipment.Weapon["off hand"];
      if (main) eq.push(`Main hand: ${main}`);
      if (off) eq.push(`Off hand: ${off}`);
    }
    if (ch.Equipment.Armor) eq.push(`Armor: ${ch.Equipment.Armor}`);
    if (ch.Equipment.Shield) eq.push(`Shield: ${ch.Equipment.Shield}`);
  }
  const eqTextarea = document.querySelector(".equipment textarea");
  if (eqTextarea) eqTextarea.value = eq.join("\n");

  const spellsTextArea = document.querySelector(".attacksandspellcasting textarea");
  if (spellsTextArea) {
    if (ch.Spellcasting && ch.Spellcasting.CanCast && Array.isArray(ch.Spellcasting.PreparedSpells)) {
      const spellList = ch.Spellcasting.PreparedSpells.map(s => `• ${s}`).join("\n");
      spellsTextArea.value = spellList || "No prepared spells.";
    } else {
      spellsTextArea.value = "No spellcasting abilities.";
    }
  }
}

function setVal(name, value) {
  const el = document.querySelector(`[name="${name}"]`);
  if (el) el.value = value ?? "";
}

function setStat(stat, score, mod) {
  setVal(`${stat}score`, score);
  const formatted = (typeof mod === "number" && mod >= 0) ? `+${mod}` : `${mod ?? ""}`;
  setVal(`${stat}mod`, formatted);
}
