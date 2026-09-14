import './css/style.css';
import './css/menu.css';
import './css/carousel.css';

import { Exit_app } from '../wailsjs/go/main/App';
import { Get_list } from '../wailsjs/go/main/App';
import { Run_file } from '../wailsjs/go/main/App';
import { SaveList } from '../wailsjs/go/main/App';
import { AddList } from '../wailsjs/go/main/App';
import { SaveSheet } from '../wailsjs/go/main/App';
import { Get_card } from '../wailsjs/go/main/App';
import { Next_card, Learn_card, Previous_card } from '../wailsjs/go/main/App';
import { Get_score } from '../wailsjs/go/main/App';
import { RemoveList } from '../wailsjs/go/main/App';
import { RidSheet } from '../wailsjs/go/main/App';
import { SwitchWhat, SwitchSelect, SwitchOrder, GetWhat } from '../wailsjs/go/main/App';

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


document.getElementById('stage').onclick = async () => {
    flip_card()
}


document.getElementById('what').onclick = async () => {
    const rev = await SwitchWhat()
    const but = document.getElementById('what');
    if (rev) {
        but.innerHTML = '<p>Show: Meaning</p>'
    } else {
        but.innerHTML = '<p>Show: Name</p>'
    }
    showFace()
}


document.getElementById('select').onclick = async () => {
    SwitchSelect()
}


document.getElementById('order').onclick = async () => {

    const shuffled = await SwitchOrder()
    const but = document.getElementById('order')

    if (shuffled) {
        but.innerHTML = '<p>Order: Random</p>'
    } else {
        but.innerHTML = '<p>Order: As Goes</p>'
    }

    const reme = await Get_card()
    await setCard(reme)
    showFace()
}


document.getElementById('previous').onclick = async () => {
    const rem = await Previous_card();
    await showFace();
    setCard(rem);
    showScore();
}

document.getElementById('learned').onclick = async () => {
    const rem = await Learn_card();
    await showFace();
    setCard(rem);
    showScore();
}

document.getElementById('next').onclick = async () => {
    const rem = await Next_card();
    await showFace();
    setCard(rem);
    showScore();
}


async function showFace() {
    const cname = document.getElementById('name');
    cname.hidden = false;
    const cp = document.getElementById('phonetic');
    cp.hidden = true;
    const cm = document.getElementById('meaning');
    cm.hidden = true;
    const cn = document.getElementById('notes');
    cn.hidden = true;

    const rev = await GetWhat ()
    if (rev) {
        flip_card()
    }
}

async function setCard(reme) {
    const cname = document.getElementById('name');
    cname.textContent = reme.Name;
    const cp = document.getElementById('phonetic');
    cp.textContent = reme.Phonetic;
    const cm = document.getElementById('meaning');
    cm.textContent = reme.Meaning;
    const cn = document.getElementById('notes');
    cn.textContent = reme.Note;
}


async function showScore() {
    const score = await Get_score();
    const board = document.getElementById('score');
    board.textContent = score.Correct + '/' + score.Reviewed + '/' + score.All;
}




function showMenu() {
    document.getElementById('menu').hidden = false
    document.getElementById('carousel').hidden = true
}

function showCarousel() {
    document.getElementById('carousel').hidden = false
    document.getElementById('menu').hidden = true
}

async function roll_file (i) {
    showCarousel();
    await Run_file(i);
    const reme = await Get_card ();
    await showFace();
    await setCard(reme);
    showScore()
}


function flip_card () {
    const cname = document.getElementById('name');
    cname.hidden = !cname.hidden;
    const cp = document.getElementById('phonetic');
    cp.hidden = !cp.hidden;
    const cm = document.getElementById('meaning');
    cm.hidden = !cm.hidden;
    const cn = document.getElementById('notes');
    cn.hidden = !cn.hidden;
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
    const btn = document.createElement('div')
    btn.className = 'run'
    btn.id = 'index' + i

    const holder = document.createElement('button')
    holder.className = 'holder'
    holder.type = 'button'

    const del = document.createElement('button')
    del.className = 'del'

    const cross = document.createElement('p')
    cross.className = 'x'
    cross.textContent = 'X'

    const nameP = document.createElement('p')
    nameP.className = 'name'
    nameP.textContent = entries[i].Name

    const pathP = document.createElement('p')
    pathP.className = 'path'
    pathP.textContent = entries[i].Path

    const cb = document.createElement('div')
    cb.className = 'cb'


    btn.appendChild(holder);
    btn.appendChild(del);
    holder.appendChild(nameP);
    holder.appendChild(pathP);
    del.appendChild(cross);
    btn.appendChild(cb);

    // capture this index for the click
    btn.addEventListener('click', async () => {
      await roll_file(i);

      const title = document.getElementById ('top');
      const h = document.createElement('h1');
      h.textContent = entries[i].Name;
      title.innerHTML = '';
      title.appendChild(h);

      render_list();
    })

    del.addEventListener('click', async (e) => {
        e.preventDefault();
        e.stopPropagation();

        await RemoveList(i);
        await render_list();
    })



    list.appendChild(btn);
  }
}



window.onload = function () {
    showMenu();
    render_list();
}

