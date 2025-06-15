import http from 'k6/http';
import { sleep } from 'k6';
import { check } from 'k6';

export let options = {
  scenarios: {
    write_and_read: {
      executor: 'constant-vus',
      vus: 5,
      duration: '20s',
      exec: 'writeAndRead',
    },
    mixed_ops: {
      executor: 'constant-vus',
      vus: 9,
      duration: '20s',
      exec: 'mixedOperations',
    },
    only_write: {
      executor: 'constant-vus',
      vus: 5,
      duration: '10s',
      exec: 'onlyWrite',
    },
    delete_random: {
      executor: 'constant-vus',
      vus: 5,
      duration: '15s',
      exec: 'deleteRandom',
      startTime: '22s',
    },
  },
};

const baseURL = 'https://localhost:8080';

/**
 * setup(): создает тестовое событие и сохраняет список всех uuid,
 * доступных для удаления, включая новый.
 */
export function setup() {
  const headers = { headers: { 'Content-Type': 'application/json' } };

  // Создаем одно событие для гарантированного uuid
  let createRes = http.post(
    `${baseURL}/event`,
    JSON.stringify({ name: 'setup-event', description: 'for deletes', date: '2025-12-12' }),
    headers
  );
  check(createRes, { 'setup POST succeeded': (r) => r.status === 200 });
  const createdUuid = createRes.json().uuid;

  // Получаем все события
  let allRes = http.get(`${baseURL}/events`);
  check(allRes, { 'GET events succeeded': (r) => r.status === 200 });
  let events = allRes.json(); // ожидаем массив объектов { uuid, ... }

  // Собираем массив uuid
  let uuids = events.map(e => e.uuid);
  if (!uuids.includes(createdUuid)) {
    uuids.push(createdUuid);
  }

  console.log('UUIDs for deletion:', uuids);
  return uuids;
}

/**
 * Сценарий: одновременно пишут и читают
 */
export function writeAndRead() {
  let res1 = http.post(
    `${baseURL}/event`,
    JSON.stringify({ name: 'event-read-write', description: 'RW', date: '2025-12-12' }),
    { headers: { 'Content-Type': 'application/json' } }
  );
  check(res1, { 'write succeeded': (r) => r.status === 200 });

  let res2 = http.get(`${baseURL}/events`);
  check(res2, { 'read succeeded': (r) => r.status === 200 });

  sleep(Math.random() * 2);
}

/**
 * Сценарий: смешанные операции (read, write, delete случайный uuid)
 */
export function mixedOperations(uuids) {
  const action = Math.random();
  if (action < 0.33) {
    http.get(`${baseURL}/events`);
  } else if (action < 0.66) {
    http.post(
      `${baseURL}/event`,
      JSON.stringify({ name: 'event-mixed', description: 'MIX', date: '2025-12-12' }),
      { headers: { 'Content-Type': 'application/json' } }
    );
  } else {
    // Выбираем случайный uuid из переданных
    const idx = Math.floor(Math.random() * uuids.length);
    const target = uuids[idx];
    http.del(`${baseURL}/event?id=${target}`);
  }
  sleep(Math.random() * 2);
}

/**
 * Сценарий: только запись
 */
export function onlyWrite() {
  let res = http.post(
    `${baseURL}/event`,
    JSON.stringify({ name: 'event-write', description: 'OnlyWrite', date: '2025-12-12' }),
    { headers: { 'Content-Type': 'application/json' } }
  );
  check(res, { 'write succeeded': (r) => r.status === 200 });
  sleep(Math.random() * 2);
}

/**
 * Сценарий: удаление случайных ресурсов из списка uuid
 */
export function deleteRandom(uuids) {
  // Берем случайный uuid
  const idx = Math.floor(Math.random() * uuids.length);
  const target = uuids[idx];
  let res = http.del(`${baseURL}/event?id=${target}`);
  check(res, { 'delete status is 200 or 404': (r) => r.status === 200 || r.status === 404 });
  sleep(Math.random() * 2);
}
