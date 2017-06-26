package main

//Catmap - These Categories have subcategories, such as Atom Physics or Machine Learning.
//They will have a separate map in order to facilitate improved error checking.
var Catmap = map[string]string{
	"stat":     "Statistics",
	"q-bio":    "Quantitative Biology",
	"cs":       "Computer Science",
	"nlin":     "Nonlinear Sciences",
	"math":     "Math",
	"cond-mat": "Physics - Mat",
	"physics":  "Physics",
}

//Primmap categories have NO SECONDARY CATEGORIES. This means that any value passed after the input will crash Arxbot
//We'll do error checking against his map when doing our Primary Category check.
var Primmap = map[string]string{
	"astro-ph": "Astrophysics",
	"gr-qc":    "General Relativity",
	"hep-ex":   "High Energy Physics - Experiment",
	"hep-lat":  "High Energy Physics - Lattice",
	"hep-ph":   "High Energy Physics - Phenomenology",
	"hep-th":   "High Energy Physics - Theory",
	"math-ph":  "Mathematical Physics",
	"nucl-ex":  "Nuclear Experiment",
	"nucl-th":  "Nuclear Theory",
	"quant-ph": "Quantum Physics",
}

//Statmap following suffix maps exist because I thought I would use them for error checking during the categories call.
//Luckily I didn't have to, but I'm loathe to delete them.
var Statmap = map[string]string{
	"AP": "Applications",
	"CO": "Computation",
	"ML": "Machine Learning",
	"ME": "Methodology",
	"TH": "Theory",
}

//Qbiomap for building search queries.
var Qbiomap = map[string]string{
	"BM": "Biomolecules",
	"CB": "Cell Behavior",
	"GN": "Genomics",
	"MN": "Molecular Networks",
	"NC": "Neurons and Cognition",
	"OT": "Other",
	"PE": "Populations and Evolution",
	"QM": "Quantitative Methods",
	"SC": "Subcellular Processes",
	"TO": "Tissues and Organs",
}

//Nlinmap for bulding search queries.
var Nlinmap = map[string]string{
	"AO": "Adaptation and Self-Organizing Systems",
	"CG": "Cellular Atuomata and Lattice Gasses",
	"CD": "Chaotic Dynamics",
	"SI": "Exactly Solvable and Integrable Systems",
	"PS": "Pattern Formation and Solitons",
}

//Mathmap for bulding search queries.
var Mathmap = map[string]string{
	"AG": "Algebraic Geometry",
	"AT": "Algebraic Topology",
	"AP": "Analysis of PDEs",
	"CT": "Category Theory",
	"CA": "Classical Analysis and ODEs",
	"CO": "Combinatorics",
	"AC": "Commutative Algebra",
	"CV": "Complex Variables",
	"DG": "Differential Geometry",
	"DS": "Dynamical Systems",
	"FA": "Functional Analysis",
	"GM": "General Mathematics",
	"GN": "General Topology",
	"GT": "Geometric Topology",
	"GR": "Group Theory",
	"HO": "History and Overview",
	"IT": "Information Theory",
	"KT": "K-Theory and Homology",
	"LO": "Logic",
	"MP": "Mathematical Physics",
	"MG": "Metric Geometry",
	"NT": "Number Theory",
	"NA": "Numerical Analysis",
	"OA": "Operator Algebras",
	"OC": "Optimization and Control",
	"PR": "Probability",
	"QA": "Quantum Algebra",
	"RT": "Representation Theory",
	"RA": "Rings and Algebras",
	"SP": "Spectral Theory",
	"ST": "Statistics",
	"SG": "Symplectic Geometry",
}

//Condmap for building search queries. 
var Condmap = map[string]string{
	"dis-nn":    "Disordered Systems and Neural Networks",
	"mes-hall":  "Mesoscopic Systems and Quantum Hall Effect",
	"mtrl-sci":  "Materials Science",
	"other":     "Other",
	"soft":      "Soft Condensed Matter",
	"stat-mech": "Statistical Mechanics",
	"str-el":    "Strongly Correlated Electrons",
	"supr-con":  "Superconductivity",
}

//Physmap for building search queries. 
var Physmap = map[string]string{
	"acc-ph":   "Accelerator Physics",
	"ao-ph":    "Atmospheric and Oceanic Physics",
	"atom-ph":  "Atomic Physics",
	"atm-clus": "Atomic and Molecular Clusters",
	"bio-ph":   "Biological Physics",
	"chem-ph":  "Chemical Physics",
	"class-ph": "Classical Physics",
	"comp-ph":  "Computational Physics",
	"data-an":  "Data Analysis; Statistics and Probability",
	"flu-dyn":  "Fluid Dynamics",
	"gen-ph":   "General Physics",
	"geo-ph":   "Geophysics",
	"hist-ph":  "History of Physics",
	"ins-det":  "Instrumentation and Detectors",
	"med-ph":   "Medical Physics",
	"optics":   "Optics",
	"ed-ph":    "Physics Education",
	"soc-ph":   "Physics and Society",
	"plasm-ph": "Plasma Physics",
	"pop-ph":   "Popular Physics",
	"space-ph": "Space Physics",
}

//CSmap for building search queries. 
var CSmap = map[string]string{
	"AR": "Architecture",
	"AI": "Artificial Intelligence",
	"CL": "Computation and Language",
	"CC": "Computational Complexity",
	"CE": "Computational Engineering; Finance; and Science",
	"CG": "Computational Geometry",
	"GT": "Computer Science and Game Theory",
	"CV": "Computer Vision and Pattern Recognition",
	"CY": "Computers and Society",
	"CR": "Cryptography and Security",
	"DS": "Data Structures and Algorithms",
	"DB": "Databases",
	"DL": "Digital Libraries",
	"DM": "Discrete Mathematics",
	"DC": "Distributed; Parallel; and Cluster Computing",
	"GL": "General Literature",
	"GR": "Graphics",
	"HC": "Human-Computer Interaction",
	"IR": "Information Retrieval",
	"IT": "Information Theory",
	"LG": "Learning",
	"LO": "Logic in Computer Science",
	"MS": "Mathematical Software",
	"MA": "Multiagent Systems",
	"MM": "Multimedia",
	"NI": "Networking and Internet Architecture",
	"NE": "Neural and Evolutionary Computing",
	"NA": "Numerical Analysis",
	"OS": "Operating Systems",
	"OH": "Other",
	"PF": "Performance",
	"PL": "Programming Languages",
	"RO": "Robotics",
	"SE": "Software Engineering",
	"SD": "Sound",
	"SC": "Symbolic Computation",
}
