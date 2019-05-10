package grid

import (
    . "github.com/DrunkenPoney/go-tictactoe/grid/tile"
    . "github.com/DrunkenPoney/go-tictactoe/position"
)

type TileGrid [][]*Tile

func NewGrid(w int, h int, tile *Tile) TileGrid {
    grid := make([][]*Tile, w)
    for x := range grid {
        grid[x] = make([]*Tile, h)
        for y := range grid[x] {
            grid[x][y] = tile.Clone()
            grid[x][y].Position = PositionAt(x, y)
        }
    }
    return grid
}

func (g TileGrid) Clone() TileGrid {
    ng := make([][]*Tile, len(g))
    for col, rows := range g {
        ng[col] = make([]*Tile, len(rows))
        for row, cell := range rows {
            ng[col][row] = cell.Clone()
        }
    }
    return ng
}

func (g TileGrid) At(pos Position) *Tile {
    x, y := pos.GetXY()
    return g[x][y]
}

func (g TileGrid) Reset() {
    for _, rows := range g {
        for _, tile := range rows {
            tile.Active = false
            tile.Winning = false
            tile.Value = EMPTY
        }
    }
}

func (g TileGrid) EmptyTiles() []*Tile {
    var empty []*Tile
    for _, rows := range g {
        for _, tile := range rows {
            if tile.Value == EMPTY {
                empty = append(empty, tile)
            }
        }
    }
    return empty
}