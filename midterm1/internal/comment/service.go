package comment

import (
	"context"
	"github.com/Tamir1205/midterm1/internal/storage/comment"
)

type service struct {
	commentRepository comment.Repository
}

type Service interface {
	AddComment(ctx context.Context, createCommentDto CreateCommentDto) error
	GetCommentsByItemId(ctx context.Context, itemId int64) ([]*Comment, error)
}

func NewService(commentRepository comment.Repository) Service {
	return &service{commentRepository: commentRepository}
}

func (s *service) AddComment(ctx context.Context, createCommentDto CreateCommentDto) error {
	_, err := s.commentRepository.CreateComment(ctx, comment.Comment{
		UserId:   createCommentDto.UserId,
		ItemId:   createCommentDto.ItemId,
		Content:  createCommentDto.Content,
		ParentId: createCommentDto.ParentId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetCommentsByItemId(ctx context.Context, itemId int64) ([]*Comment, error) {
	comments, err := s.commentRepository.FindCommentByItemId(ctx, itemId)
	if err != nil {
		return nil, err
	}

	dtoComments := make([]*Comment, len(comments))
	for i, c := range comments {
		dtoComments[i] = repoToDtoComment(c)
	}

	dtoComments = buildHierarchy(dtoComments)

	return dtoComments, err
}

func repoToDtoComment(comment comment.Comment) *Comment {
	return &Comment{
		ID:       comment.ID,
		UserId:   comment.UserId,
		ItemId:   comment.ItemId,
		Content:  comment.Content,
		ParentId: comment.ParentId,
	}
}

func buildHierarchy(comments []*Comment) []*Comment {
	commentMap := make(map[int64]*Comment)
	roots := make([]*Comment, 0)

	for _, comment := range comments {
		commentMap[comment.ID] = comment
		if comment.ParentId == nil {
			roots = append(roots, comment)
		}
	}

	for _, comment := range comments {
		if comment.ParentId != nil {
			parent := commentMap[*comment.ParentId]
			parent.Children = append(parent.Children, *comment)
		}
	}

	return roots
}
