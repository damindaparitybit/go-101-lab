// add speakers here and run : npm run prepare to make go-200 in sync with go-100
const speakers = [
  // comment speakers before presenting
  'SFR'
].map((trigram) => `speakers/${trigram}`);

// go100 : canonical structure for ordered 2-levels filesystem tree
// array of [directory, [file]] pairs
const go100 = [
  ['00-school', ['00-title', ...speakers]],
  ['01-intro', ['00-title', '01-go_purpose']],
  ['02-installation', ['00-titles', '01-env_install']],
  ['03-let_s_go', ['00-titles', '01-basics', '02-if', '03-loops']],
  [
    '04-lets_go_further',
    ['00-title', '01-types_structures', '02-pointers', '03-arrays_slices_map'],
  ],
  ['05-object_oriented', ['00-title', '01-methods', '02-interfaces']],
  [
    '06-concurrency',
    [
      '00-title',
      '01-goroutine',
      '02-channels',
      '03-buffered-channels',
      '04-closing-channels',
      '05-range',
      '06-select',
      '07-exercise',
    ],
  ],
  ['07-and_now', ['00-title', '01-going_further', '02-thanks']],
];

const makeSlidePath = (dir) => (file) => ({ path: `${dir}/${file}.md` });
const pathReducerFactory = (baseDir = '') => (acc, [dir, files]) => [
  ...acc,
  ...files.map(makeSlidePath(`${baseDir}${dir}`)),
];

const makeSlidePaths = ([slides, baseDir]) => slides.reduce(pathReducerFactory(baseDir), []);

// comment the presentation you don't want to use
export const slides = [
  [go100, 'go-100/'],
]
  .map(makeSlidePaths)
  .flat();
