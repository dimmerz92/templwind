package themes

type TailwindColor string

const (
	Black TailwindColor = "black"
	White TailwindColor = "white"

	Current     TailwindColor = "currentcolor"
	Transparent TailwindColor = "transparent"

	Slate50  TailwindColor = "var(--color-slate-50)"
	Slate100 TailwindColor = "var(--color-slate-100)"
	Slate200 TailwindColor = "var(--color-slate-200)"
	Slate300 TailwindColor = "var(--color-slate-300)"
	Slate400 TailwindColor = "var(--color-slate-400)"
	Slate500 TailwindColor = "var(--color-slate-500)"
	Slate600 TailwindColor = "var(--color-slate-600)"
	Slate700 TailwindColor = "var(--color-slate-700)"
	Slate800 TailwindColor = "var(--color-slate-800)"
	Slate900 TailwindColor = "var(--color-slate-900)"
	Slate950 TailwindColor = "var(--color-slate-950)"

	Gray50  TailwindColor = "var(--color-gray-50)"
	Gray100 TailwindColor = "var(--color-gray-100)"
	Gray200 TailwindColor = "var(--color-gray-200)"
	Gray300 TailwindColor = "var(--color-gray-300)"
	Gray400 TailwindColor = "var(--color-gray-400)"
	Gray500 TailwindColor = "var(--color-gray-500)"
	Gray600 TailwindColor = "var(--color-gray-600)"
	Gray700 TailwindColor = "var(--color-gray-700)"
	Gray800 TailwindColor = "var(--color-gray-800)"
	Gray900 TailwindColor = "var(--color-gray-900)"
	Gray950 TailwindColor = "var(--color-gray-950)"

	Zinc50  TailwindColor = "var(--color-zinc-50)"
	Zinc100 TailwindColor = "var(--color-zinc-100)"
	Zinc200 TailwindColor = "var(--color-zinc-200)"
	Zinc300 TailwindColor = "var(--color-zinc-300)"
	Zinc400 TailwindColor = "var(--color-zinc-400)"
	Zinc500 TailwindColor = "var(--color-zinc-500)"
	Zinc600 TailwindColor = "var(--color-zinc-600)"
	Zinc700 TailwindColor = "var(--color-zinc-700)"
	Zinc800 TailwindColor = "var(--color-zinc-800)"
	Zinc900 TailwindColor = "var(--color-zinc-900)"
	Zinc950 TailwindColor = "var(--color-zinc-950)"

	Neutral50  TailwindColor = "var(--color-neutral-50)"
	Neutral100 TailwindColor = "var(--color-neutral-100)"
	Neutral200 TailwindColor = "var(--color-neutral-200)"
	Neutral300 TailwindColor = "var(--color-neutral-300)"
	Neutral400 TailwindColor = "var(--color-neutral-400)"
	Neutral500 TailwindColor = "var(--color-neutral-500)"
	Neutral600 TailwindColor = "var(--color-neutral-600)"
	Neutral700 TailwindColor = "var(--color-neutral-700)"
	Neutral800 TailwindColor = "var(--color-neutral-800)"
	Neutral900 TailwindColor = "var(--color-neutral-900)"
	Neutral950 TailwindColor = "var(--color-neutral-950)"

	Stone50  TailwindColor = "var(--color-stone-50)"
	Stone100 TailwindColor = "var(--color-stone-100)"
	Stone200 TailwindColor = "var(--color-stone-200)"
	Stone300 TailwindColor = "var(--color-stone-300)"
	Stone400 TailwindColor = "var(--color-stone-400)"
	Stone500 TailwindColor = "var(--color-stone-500)"
	Stone600 TailwindColor = "var(--color-stone-600)"
	Stone700 TailwindColor = "var(--color-stone-700)"
	Stone800 TailwindColor = "var(--color-stone-800)"
	Stone900 TailwindColor = "var(--color-stone-900)"
	Stone950 TailwindColor = "var(--color-stone-950)"

	Red50  TailwindColor = "var(--color-red-50)"
	Red100 TailwindColor = "var(--color-red-100)"
	Red200 TailwindColor = "var(--color-red-200)"
	Red300 TailwindColor = "var(--color-red-300)"
	Red400 TailwindColor = "var(--color-red-400)"
	Red500 TailwindColor = "var(--color-red-500)"
	Red600 TailwindColor = "var(--color-red-600)"
	Red700 TailwindColor = "var(--color-red-700)"
	Red800 TailwindColor = "var(--color-red-800)"
	Red900 TailwindColor = "var(--color-red-900)"
	Red950 TailwindColor = "var(--color-red-950)"

	Orange50  TailwindColor = "var(--color-orange-50)"
	Orange100 TailwindColor = "var(--color-orange-100)"
	Orange200 TailwindColor = "var(--color-orange-200)"
	Orange300 TailwindColor = "var(--color-orange-300)"
	Orange400 TailwindColor = "var(--color-orange-400)"
	Orange500 TailwindColor = "var(--color-orange-500)"
	Orange600 TailwindColor = "var(--color-orange-600)"
	Orange700 TailwindColor = "var(--color-orange-700)"
	Orange800 TailwindColor = "var(--color-orange-800)"
	Orange900 TailwindColor = "var(--color-orange-900)"
	Orange950 TailwindColor = "var(--color-orange-950)"

	Amber50  TailwindColor = "var(--color-amber-50)"
	Amber100 TailwindColor = "var(--color-amber-100)"
	Amber200 TailwindColor = "var(--color-amber-200)"
	Amber300 TailwindColor = "var(--color-amber-300)"
	Amber400 TailwindColor = "var(--color-amber-400)"
	Amber500 TailwindColor = "var(--color-amber-500)"
	Amber600 TailwindColor = "var(--color-amber-600)"
	Amber700 TailwindColor = "var(--color-amber-700)"
	Amber800 TailwindColor = "var(--color-amber-800)"
	Amber900 TailwindColor = "var(--color-amber-900)"
	Amber950 TailwindColor = "var(--color-amber-950)"

	Yellow50  TailwindColor = "var(--color-yellow-50)"
	Yellow100 TailwindColor = "var(--color-yellow-100)"
	Yellow200 TailwindColor = "var(--color-yellow-200)"
	Yellow300 TailwindColor = "var(--color-yellow-300)"
	Yellow400 TailwindColor = "var(--color-yellow-400)"
	Yellow500 TailwindColor = "var(--color-yellow-500)"
	Yellow600 TailwindColor = "var(--color-yellow-600)"
	Yellow700 TailwindColor = "var(--color-yellow-700)"
	Yellow800 TailwindColor = "var(--color-yellow-800)"
	Yellow900 TailwindColor = "var(--color-yellow-900)"
	Yellow950 TailwindColor = "var(--color-yellow-950)"

	Lime50  TailwindColor = "var(--color-lime-50)"
	Lime100 TailwindColor = "var(--color-lime-100)"
	Lime200 TailwindColor = "var(--color-lime-200)"
	Lime300 TailwindColor = "var(--color-lime-300)"
	Lime400 TailwindColor = "var(--color-lime-400)"
	Lime500 TailwindColor = "var(--color-lime-500)"
	Lime600 TailwindColor = "var(--color-lime-600)"
	Lime700 TailwindColor = "var(--color-lime-700)"
	Lime800 TailwindColor = "var(--color-lime-800)"
	Lime900 TailwindColor = "var(--color-lime-900)"
	Lime950 TailwindColor = "var(--color-lime-950)"

	Green50  TailwindColor = "var(--color-green-50)"
	Green100 TailwindColor = "var(--color-green-100)"
	Green200 TailwindColor = "var(--color-green-200)"
	Green300 TailwindColor = "var(--color-green-300)"
	Green400 TailwindColor = "var(--color-green-400)"
	Green500 TailwindColor = "var(--color-green-500)"
	Green600 TailwindColor = "var(--color-green-600)"
	Green700 TailwindColor = "var(--color-green-700)"
	Green800 TailwindColor = "var(--color-green-800)"
	Green900 TailwindColor = "var(--color-green-900)"
	Green950 TailwindColor = "var(--color-green-950)"

	Emerald50  TailwindColor = "var(--color-emerald-50)"
	Emerald100 TailwindColor = "var(--color-emerald-100)"
	Emerald200 TailwindColor = "var(--color-emerald-200)"
	Emerald300 TailwindColor = "var(--color-emerald-300)"
	Emerald400 TailwindColor = "var(--color-emerald-400)"
	Emerald500 TailwindColor = "var(--color-emerald-500)"
	Emerald600 TailwindColor = "var(--color-emerald-600)"
	Emerald700 TailwindColor = "var(--color-emerald-700)"
	Emerald800 TailwindColor = "var(--color-emerald-800)"
	Emerald900 TailwindColor = "var(--color-emerald-900)"
	Emerald950 TailwindColor = "var(--color-emerald-950)"

	Teal50  TailwindColor = "var(--color-teal-50)"
	Teal100 TailwindColor = "var(--color-teal-100)"
	Teal200 TailwindColor = "var(--color-teal-200)"
	Teal300 TailwindColor = "var(--color-teal-300)"
	Teal400 TailwindColor = "var(--color-teal-400)"
	Teal500 TailwindColor = "var(--color-teal-500)"
	Teal600 TailwindColor = "var(--color-teal-600)"
	Teal700 TailwindColor = "var(--color-teal-700)"
	Teal800 TailwindColor = "var(--color-teal-800)"
	Teal900 TailwindColor = "var(--color-teal-900)"
	Teal950 TailwindColor = "var(--color-teal-950)"

	Cyan50  TailwindColor = "var(--color-cyan-50)"
	Cyan100 TailwindColor = "var(--color-cyan-100)"
	Cyan200 TailwindColor = "var(--color-cyan-200)"
	Cyan300 TailwindColor = "var(--color-cyan-300)"
	Cyan400 TailwindColor = "var(--color-cyan-400)"
	Cyan500 TailwindColor = "var(--color-cyan-500)"
	Cyan600 TailwindColor = "var(--color-cyan-600)"
	Cyan700 TailwindColor = "var(--color-cyan-700)"
	Cyan800 TailwindColor = "var(--color-cyan-800)"
	Cyan900 TailwindColor = "var(--color-cyan-900)"
	Cyan950 TailwindColor = "var(--color-cyan-950)"

	Sky50  TailwindColor = "var(--color-sky-50)"
	Sky100 TailwindColor = "var(--color-sky-100)"
	Sky200 TailwindColor = "var(--color-sky-200)"
	Sky300 TailwindColor = "var(--color-sky-300)"
	Sky400 TailwindColor = "var(--color-sky-400)"
	Sky500 TailwindColor = "var(--color-sky-500)"
	Sky600 TailwindColor = "var(--color-sky-600)"
	Sky700 TailwindColor = "var(--color-sky-700)"
	Sky800 TailwindColor = "var(--color-sky-800)"
	Sky900 TailwindColor = "var(--color-sky-900)"
	Sky950 TailwindColor = "var(--color-sky-950)"

	Blue50  TailwindColor = "var(--color-blue-50)"
	Blue100 TailwindColor = "var(--color-blue-100)"
	Blue200 TailwindColor = "var(--color-blue-200)"
	Blue300 TailwindColor = "var(--color-blue-300)"
	Blue400 TailwindColor = "var(--color-blue-400)"
	Blue500 TailwindColor = "var(--color-blue-500)"
	Blue600 TailwindColor = "var(--color-blue-600)"
	Blue700 TailwindColor = "var(--color-blue-700)"
	Blue800 TailwindColor = "var(--color-blue-800)"
	Blue900 TailwindColor = "var(--color-blue-900)"
	Blue950 TailwindColor = "var(--color-blue-950)"

	Indigo50  TailwindColor = "var(--color-indigo-50)"
	Indigo100 TailwindColor = "var(--color-indigo-100)"
	Indigo200 TailwindColor = "var(--color-indigo-200)"
	Indigo300 TailwindColor = "var(--color-indigo-300)"
	Indigo400 TailwindColor = "var(--color-indigo-400)"
	Indigo500 TailwindColor = "var(--color-indigo-500)"
	Indigo600 TailwindColor = "var(--color-indigo-600)"
	Indigo700 TailwindColor = "var(--color-indigo-700)"
	Indigo800 TailwindColor = "var(--color-indigo-800)"
	Indigo900 TailwindColor = "var(--color-indigo-900)"
	Indigo950 TailwindColor = "var(--color-indigo-950)"

	Violet50  TailwindColor = "var(--color-violet-50)"
	Violet100 TailwindColor = "var(--color-violet-100)"
	Violet200 TailwindColor = "var(--color-violet-200)"
	Violet300 TailwindColor = "var(--color-violet-300)"
	Violet400 TailwindColor = "var(--color-violet-400)"
	Violet500 TailwindColor = "var(--color-violet-500)"
	Violet600 TailwindColor = "var(--color-violet-600)"
	Violet700 TailwindColor = "var(--color-violet-700)"
	Violet800 TailwindColor = "var(--color-violet-800)"
	Violet900 TailwindColor = "var(--color-violet-900)"
	Violet950 TailwindColor = "var(--color-violet-950)"

	Purple50  TailwindColor = "var(--color-purple-50)"
	Purple100 TailwindColor = "var(--color-purple-100)"
	Purple200 TailwindColor = "var(--color-purple-200)"
	Purple300 TailwindColor = "var(--color-purple-300)"
	Purple400 TailwindColor = "var(--color-purple-400)"
	Purple500 TailwindColor = "var(--color-purple-500)"
	Purple600 TailwindColor = "var(--color-purple-600)"
	Purple700 TailwindColor = "var(--color-purple-700)"
	Purple800 TailwindColor = "var(--color-purple-800)"
	Purple900 TailwindColor = "var(--color-purple-900)"
	Purple950 TailwindColor = "var(--color-purple-950)"

	Fuchsia50  TailwindColor = "var(--color-fuchsia-50)"
	Fuchsia100 TailwindColor = "var(--color-fuchsia-100)"
	Fuchsia200 TailwindColor = "var(--color-fuchsia-200)"
	Fuchsia300 TailwindColor = "var(--color-fuchsia-300)"
	Fuchsia400 TailwindColor = "var(--color-fuchsia-400)"
	Fuchsia500 TailwindColor = "var(--color-fuchsia-500)"
	Fuchsia600 TailwindColor = "var(--color-fuchsia-600)"
	Fuchsia700 TailwindColor = "var(--color-fuchsia-700)"
	Fuchsia800 TailwindColor = "var(--color-fuchsia-800)"
	Fuchsia900 TailwindColor = "var(--color-fuchsia-900)"
	Fuchsia950 TailwindColor = "var(--color-fuchsia-950)"

	Pink50  TailwindColor = "var(--color-pink-50)"
	Pink100 TailwindColor = "var(--color-pink-100)"
	Pink200 TailwindColor = "var(--color-pink-200)"
	Pink300 TailwindColor = "var(--color-pink-300)"
	Pink400 TailwindColor = "var(--color-pink-400)"
	Pink500 TailwindColor = "var(--color-pink-500)"
	Pink600 TailwindColor = "var(--color-pink-600)"
	Pink700 TailwindColor = "var(--color-pink-700)"
	Pink800 TailwindColor = "var(--color-pink-800)"
	Pink900 TailwindColor = "var(--color-pink-900)"
	Pink950 TailwindColor = "var(--color-pink-950)"

	Rose50  TailwindColor = "var(--color-rose-50)"
	Rose100 TailwindColor = "var(--color-rose-100)"
	Rose200 TailwindColor = "var(--color-rose-200)"
	Rose300 TailwindColor = "var(--color-rose-300)"
	Rose400 TailwindColor = "var(--color-rose-400)"
	Rose500 TailwindColor = "var(--color-rose-500)"
	Rose600 TailwindColor = "var(--color-rose-600)"
	Rose700 TailwindColor = "var(--color-rose-700)"
	Rose800 TailwindColor = "var(--color-rose-800)"
	Rose900 TailwindColor = "var(--color-rose-900)"
	Rose950 TailwindColor = "var(--color-rose-950)"
)