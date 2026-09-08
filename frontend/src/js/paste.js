// includes.js — RECURSIVE VERSION (based exactly on your working code)

document.addEventListener('DOMContentLoaded', () => {
  loadIncludesRecursively();
});

async function loadIncludesRecursively(container = document.documentElement) {
  const elements = container.querySelectorAll('[data-include]:not([data-include-loaded])');

  for (const element of elements) {
    const file = element.getAttribute('data-include');
    
    // Mark as processed so we don't create loops
    element.setAttribute('data-include-loaded', 'true');

    try {
      const response = await fetch(file);
      if (!response.ok) throw new Error(`Failed to load ${file}`);
      
      const html = await response.text();
      element.innerHTML = html;

      // ← RECURSIVE MAGIC: load any new data-include elements inside what we just inserted
      await loadIncludesRecursively(element);
    } catch (error) {
      console.error(error);
      element.innerHTML = `<p style="color:red; float:right; margin-top: 30px;">⚠️ Could not load sidebar</p>`;
    }
  }
}
