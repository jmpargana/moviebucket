import { writable, derived } from 'svelte/store'
import { movies } from '../api/movies'

function createPersonalMovieBucket() {
  const { subscribe, update } = writable(movies)
  return {
    subscribe,
    add: (movie) => update(movies => [...movies, movie])
  }
}

export const personalMovieBucket = createPersonalMovieBucket()

export const dailyPick = derived(personalMovieBucket, $movies => $movies[Math.round(Math.random() * $movies.length)])
