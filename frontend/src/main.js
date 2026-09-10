import './css/style.css';
import './css/menu.css';
import './css/carousel.css'

import { Exit_app } from '../wailsjs/go/main/App';
import { Get_list } from '../wailsjs/go/main/App';
import { Run_file } from '../wailsjs/go/main/App';
import { SaveList } from '../wailsjs/go/main/App';
import { AddList } from '../wailsjs/go/main/App';
import { SaveSheet } from '../wailsjs/go/main/App';

document.addEventListener('keydown', (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'q') {
        e.preventDefault();
        Exit_app();
    }
})







document.getElementById('save_list').onclick = async () => {
    SaveList ()
}


document.getElementById('add_new').onclick = async () => {
    await AddList ()
    await render_list()
    await Run_file (0)
}


document.getElementById('go_to_menu').onclick = async () => {
    showMenu()
}



document.getElementById('save_file').onclick = async () => {
    SaveSheet()
}




function showMenu() {
    const stage = document.getElementById('stage')
    stage.innerHTML = '<p id="phonetic">Loading...</p>'
    const top = document.getElementById('top')
    top.innerHTML = '<h1 style="color: #333333;">Loading...</h1>'
    document.getElementById('menu').hidden = false
    document.getElementById('carousel').hidden = true
}

function showCarousel() {
    document.getElementById('carousel').hidden = false
    document.getElementById('menu').hidden = true
}



async function render_list() {
  const list = document.getElementById('list')
  list.innerHTML = '' // clear

  const entries = await Get_list()
  if (!entries || entries.length === 0) {
    list.innerHTML = '<p class="def">There are no files yet.</p>'
    return
  }


    for (let i = 0; i < entries.length; i++) {
    const btn = document.createElement('button')
    btn.className = 'run'
    btn.id = 'index' + i

    const nameP = document.createElement('p')
    nameP.className = 'name'
    nameP.textContent = entries[i].Name

    const pathP = document.createElement('p')
    pathP.className = 'path'
    pathP.textContent = entries[i].Path

    btn.appendChild(nameP)
    btn.appendChild(pathP)

    // capture this index for the click
    btn.addEventListener('click', () => {
      showCarousel()
      Run_file(i)

      const title = document.getElementById ('top')
      const h = document.createElement('h1')
      h.textContent = entries[i].Name
      title.innerHTML = ''
      title.appendChild(h)

      render_list()
    })

    list.appendChild(btn)
  }
}



window.onload = function () {
    showMenu();
    render_list();
}

