# Sphinx Doc
## Installieren

```
pip install -U sphinx
pip install -U myst_parser

```
## Neues Dokument erstellen
```

sphinx-quickstart


```

## config.py

```

extensions = ['myst_parser']
source_suffix = {
    '.rst': 'restructuredtext',
    '.txt': 'markdown',
    '.md': 'markdown',
}

```



