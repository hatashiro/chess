import {chess} from './enums.js';

const pieceMap = {
  [chess.Player.P1]: {
    [chess.PieceType.KING]:   '♚',
    [chess.PieceType.QUEEN]:  '♛',
    [chess.PieceType.ROOK]:   '♜',
    [chess.PieceType.BISHOP]: '♝',
    [chess.PieceType.KNIGHT]: '♞',
    [chess.PieceType.PAWN]:   '♟',
  },
  [chess.Player.P2]: {
    [chess.PieceType.KING]:   '♔',
    [chess.PieceType.QUEEN]:  '♕',
    [chess.PieceType.ROOK]:   '♖',
    [chess.PieceType.BISHOP]: '♗',
    [chess.PieceType.KNIGHT]: '♘',
    [chess.PieceType.PAWN]:   '♙',
  },
};

export const piece = ({player, type}) => pieceMap[player][type];
