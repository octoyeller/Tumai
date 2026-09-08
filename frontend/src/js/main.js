import '../css/style.css';
import '../css/app.css';

import { Exit_app } from '../../wailsjs/go/main/App';
import { Get_list } from '../../wailsjs/go/main/App';

document.addEventListener('keydown', (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'q') {
        e.preventDefault();
        Exit_app();
    }
})







document.getElementById('save').onclick = async () => {
    console.log('save')
}


document.getElementById('add_new').onclick = async () => {
    console.log('new')
}



async function showFirstEntry() {
  const list = document.getElementById('list')
  list.innerHTML = '' // clear

  const entries = await Get_list()
  if (!entries || entries.length === 0) {
    list.textContent = 'No entries'
    return
  }

  const entry = entries[0]

  const btn = document.createElement('button')
  btn.className = 'run'
  btn.id = 'index0'

  const nameP = document.createElement('p')
  nameP.className = 'name'
  nameP.textContent = entry.Name   // from Go struct

  const pathP = document.createElement('p')
  pathP.className = 'path'
  pathP.textContent = entry.Path

  btn.appendChild(nameP)
  btn.appendChild(pathP)
  list.appendChild(btn)
}


window.onload = function () {
    showFirstEntry();
}

