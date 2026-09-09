import '../css/style.css';
import '../css/app.css';

import { Exit_app } from '../../wailsjs/go/main/App';
import { Get_list } from '../../wailsjs/go/main/App';
import { Run_file } from '../../wailsjs/go/main/App';
import { SaveList } from '../../wailsjs/go/main/App';

document.addEventListener('keydown', (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'q') {
        e.preventDefault();
        Exit_app();
    }
})







document.getElementById('save').onclick = async () => {
    SaveList ()
}


document.getElementById('add_new').onclick = async () => {
    console.log('new')
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
      Run_file(i)
    })

    list.appendChild(btn)
  }
}



window.onload = function () {
    render_list();
}

