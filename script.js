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
    const res = await fetch("/data/characters.json");
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

function fillSkills(ch) {
  if (!ch.Skills || !ch.Stats) return;

  const skillMap = {
    "Acrobatics": "DexMod",
    "Animal Handling": "WisMod",
    "Arcana": "IntelMod",
    "Athletics": "StrMod",
    "Deception": "ChaMod",
    "History": "IntelMod",
    "Insight": "WisMod",
    "Intimidation": "ChaMod",
    "Investigation": "IntelMod",
    "Medicine": "WisMod",
    "Nature": "IntelMod",
    "Perception": "WisMod",
    "Performance": "ChaMod",
    "Persuasion": "ChaMod",
    "Religion": "IntelMod",
    "Sleight of Hand": "DexMod",
    "Stealth": "DexMod",
    "Survival": "WisMod"
  };

  const profSkills = (ch.Skills.Skills || []).map(s => s.toLowerCase());
  const profBonus = ch.Proficiency ?? 0;

  for (const [skill, abilityModKey] of Object.entries(skillMap)) {
    const isProf = profSkills.includes(skill.toLowerCase());
    const baseMod = ch.Stats?.[abilityModKey] ?? 0;
    const total = baseMod + (isProf ? profBonus : 0);

    const field = document.querySelector(`[name="${skill}"]`);
    if (field) field.value = total >= 0 ? `+${total}` : `${total}`;

    const checkbox = document.querySelector(`[name="${skill}-prof"]`);
    if (checkbox) checkbox.checked = isProf;
  }
}

const classSaveProficiencies = {
  barbarian: ["strength", "constitution"],
  bard: ["dexterity", "charisma"],
  cleric: ["wisdom", "charisma"],
  druid: ["intelligence", "wisdom"],
  fighter: ["strength", "constitution"],
  monk: ["strength", "dexterity"],
  paladin: ["wisdom", "charisma"],
  ranger: ["strength", "dexterity"],
  rogue: ["dexterity", "intelligence"],
  sorcerer: ["constitution", "charisma"],
  warlock: ["wisdom", "charisma"],
  wizard: ["intelligence", "wisdom"]
};

function fillSavingThrows(ch) {
  if (!ch.Stats || !ch.Class) return;

  const saves = {
    Strength: ch.Stats.StrMod,
    Dexterity: ch.Stats.DexMod,
    Constitution: ch.Stats.ConMod,
    Intelligence: ch.Stats.IntelMod,
    Wisdom: ch.Stats.WisMod,
    Charisma: ch.Stats.ChaMod
  };

  const profBonus = ch.Proficiency ?? 0;
  const profSaves = classSaveProficiencies[ch.Class.toLowerCase()] || [];

  for (const [ability, baseMod] of Object.entries(saves)) {
    const isProf = profSaves.includes(ability.toLowerCase());
    const total = baseMod + (isProf ? profBonus : 0);

    const field = document.querySelector(`[name="${ability}-save"]`);
    if (field) field.value = total >= 0 ? `+${total}` : `${total}`;

    const checkbox = document.querySelector(`[name="${ability}-save-prof"]`);
    if (checkbox) checkbox.checked = isProf;
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

  const attackRows = document.querySelectorAll(".attacksandspellcasting table tbody tr");
  attackRows.forEach(row => row.querySelectorAll("input").forEach(i => i.value = ""));

  const weaponEntries = [];
  if (ch.Equipment?.Weapon) {
    const main = ch.Equipment.Weapon["main hand"];
    const off = ch.Equipment.Weapon["off hand"];
    if (main) weaponEntries.push(main);
    if (off) weaponEntries.push(off);
  }

  const sc = ch.Spellcasting || {};
  const spells = (sc.prepared_spells?.length ? sc.prepared_spells : sc.learned_spells) ?? [];

  const spellsTextArea = document.querySelector(".attacksandspellcasting textarea");

  const combinedAttacks = [
    ...weaponEntries.map(w => ({ type: "weapon", name: w })),
    ...spells.slice(0, 3).map(s => ({ type: "spell", name: s })),
  ];

  combinedAttacks.slice(0, 3).forEach((atk, i) => {
    const row = attackRows[i];
    if (!row) return;

    const [nameField, atkField, dmgField] = row.querySelectorAll("input");

    if (atk.type === "weapon") {
      const dexMod = ch.Stats?.DexMod ?? 0;
      const strMod = ch.Stats?.StrMod ?? 0;
      const prof = ch.Proficiency ?? 0;
      const name = atk.name?.toLowerCase() ?? "";
      const isFinesse = ["rapier", "shortsword", "dagger"].includes(name);
      const isRanged = name.includes("bow") || name.includes("crossbow") || name.includes("sling");

      const mod = isFinesse ? Math.max(dexMod, strMod) : (isRanged ? dexMod : strMod);
      const totalBonus = mod + prof;

      nameField.value = atk.name;
      atkField.value = totalBonus >= 0 ? `+${totalBonus}` : `${totalBonus}`;
    }

    if (atk.type === "spell") {
      const bonus = sc.SpellAttackBonus ?? 0;
      nameField.value = atk.name;
      atkField.value = bonus >= 0 ? `+${bonus}` : `${bonus}`;
      dmgField.value = "—";
    }
  });

  fillSkills(ch);
  fillSavingThrows(ch);
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
